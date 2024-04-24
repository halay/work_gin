package yto

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const key = "justcall"

func DownloadMP3Handler(c *gin.Context) {
	var params struct {
		Ts         int64  `json:"ts" form:"ts"`
		BranchCode string `json:"branch_code" form:"branch_code"`
		CallID     string `json:"call_id" form:"call_id"`
		CallFrom   string `json:"call_from" form:"call_from"`
	}
	Token := c.GetHeader("Token")
	if Token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Token is required"})
		return
	}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid request parameters"})
		return
	}
	// Check timestamp validity (within the last 5 minutes)
	now := time.Now().Unix()
	if now-params.Ts > 5*60 { // 5 minutes in seconds
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Request timestamp is expired"})
		return
	}

	// Verify signature
	expectedSign := generateSignature(params.Ts, params.BranchCode, params.CallID, params.CallFrom)
	if Token != expectedSign {
		c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Invalid signature"})
		return
	}
	// Extract year, month, day from call_id
	callIDParts := strings.Split(params.CallID, ".")
	timeStampStr := callIDParts[0]

	timestamp, err := strconv.ParseInt(timeStampStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid call_id format"})
		return
	}

	callTimestamp := time.Unix(timestamp, 0)
	year, month, day := callTimestamp.Date()
	var dayStr, monthStr string
	if day < 10 {
		dayStr = "0" + strconv.Itoa(day)
	} else {
		dayStr = strconv.Itoa(day)
	}
	if month < 10 {
		monthStr = "0" + strconv.Itoa(int(month))
	} else {
		monthStr = strconv.Itoa(int(month))
	}
	// Assuming the file path is constructed using branch_code and call_id
	var filePath string
	if params.CallFrom == "3" { //云CC
		filePath = fmt.Sprintf("/monitor/ycc/%s/%d/%s/%s/%s.mp3", params.BranchCode, year, monthStr, dayStr, params.CallID)
	} else if params.CallFrom == "1" { //总部
		yearAfterTwoDigits := strconv.Itoa(year % 100)
		filePath = fmt.Sprintf("/monitor/zongbu/%s-%s/%s/%s.mp3", yearAfterTwoDigits, monthStr, dayStr, params.CallID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "Invalid call_from"})
		return
	}
	// Check file existence and permissions
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": "File not found" + filePath})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"code": http.StatusInternalServerError, "message": "Internal server error"})
		}
		return
	}

	// Set HTTP response headers for file download
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileInfo.Name()))
	c.Header("Content-Type", "audio/mpeg")
	c.Header("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

	// Serve the MP3 file
	http.ServeFile(c.Writer, c.Request, filePath)
}

func generateSignature(ts int64, branchCode, callID, callFrom string) string {
	data := fmt.Sprintf("%d%s%s%s%s", ts, branchCode, callID, callFrom, key)
	hash := md5.Sum([]byte(data))
	return hex.EncodeToString(hash[:])
}
