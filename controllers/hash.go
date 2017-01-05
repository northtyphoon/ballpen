package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/northtyphoon/ballpen/uuid"
)

// HashRequest represents a hash request payload
type HashRequest struct {
	Algorithm  string
	OriContent string
}

// HashController is the hash API controller
type HashController struct {
	beego.Controller
}

// Post verb
func (c *HashController) Post() {
	var request HashRequest

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &request); err == nil && request.Algorithm != "" {
		if hashedContent, err := hash(request.Algorithm, request.OriContent); err == nil {
			c.Data["json"] = map[string]interface{}{"hashedContent": hashedContent}
		} else {
			logs.Debug("Hash.Post.Request: %v", request)
			c.Ctx.Output.SetStatus(400)
			c.Data["json"] = map[string]interface{}{"message": err.Error()}
		}
	} else {
		logs.Debug("Hash.Post.Request.Raw: %x", c.Ctx.Input.RequestBody)
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = map[string]interface{}{"message": err.Error()}
	}

	c.ServeJSON()
}

func hash(algorithm string, oriContent string) (string, error) {
	if algorithm == "sha256" {
		h256 := sha256.New()
		h256.Write(transformToUUID(oriContent))
		return fmt.Sprintf("%x", h256.Sum(nil)), nil
	} else if algorithm == "md5" {
		md5 := md5.New()
		md5.Write(transformToUUID(oriContent))
		return fmt.Sprintf("%x", md5.Sum(nil)), nil
	}

	err := fmt.Errorf("The hash algorithm %s is not supported", algorithm)
	return "", err
}

func transformToUUID(oriContent string) []byte {
	if u, err := uuid.Parse(oriContent); err == nil {
		return u[:]
	}

	return []byte(oriContent)
}
