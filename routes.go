package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hackthenorth2020/go-firebase/profiles"
)

func createProfile(c *gin.Context) {
	var createProfileRequest *profiles.Profile
	err := c.Bind(&createProfileRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := profileSrv.CreateProfile(createProfileRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := Atoi(idStr)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := profileSrv.ReadProfile(&id)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func updateProfile(c *gin.Context) {
	var updateProfileRequest *profiles.Profile
	err := c.Bind(&updateProfileRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := profileSrv.UpdateItem(updateProfileRequest)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := Atoi(idStr)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	resp, err := profileSrv.DeleteProfile(&id)
	if err != nil {
		c.AbortWithError(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readAllProfiles(c *gin.Context) {

	resp, err := profileSrv.ReadAllProfiles()
		return
	}

	c.JSON(200, &resp)

}
