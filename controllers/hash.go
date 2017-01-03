package controllers

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type HashRequest struct {
	Algorithm  string
	OriContent string
}

type HashController struct {
	beego.Controller
}

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

func hash(algortihm string, oriContent string) (string, error) {
	if algortihm == "sha256" {
		h256 := sha256.New()
		h256.Write([]byte(oriContent))
		return base64.StdEncoding.EncodeToString(h256.Sum(nil)), nil
	} else if algortihm == "md5" {
		md5 := md5.New()
		md5.Write([]byte(oriContent))
		return base64.StdEncoding.EncodeToString(md5.Sum(nil)), nil
	}

	err := fmt.Errorf("The hash algorithm %s is not supported", algortihm)
	return "", err
}
