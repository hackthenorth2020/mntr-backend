package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hackthenorth2020/go-firebase/profiles"
)

func createProfile(c *gin.Context) {
	var createProfileRequest *profiles.Profile
	err := c.Bind(&createProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	if createProfileRequest.UID == nil {
		createProfileRequest.UID = c.GetString("UID")
	} else if createProfileRequest.UID != c.GetString("UID") {
		c.JSON("402", gin.H{"error", "token UID does not match request UID"})
	}

	resp, err := profileSrv.CreateProfile(createProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readProfile(c *gin.Context) {
	idStr := c.Param("id")
	id, err := Atoi(idStr)
	if err != nil {
		c.JSON(501, err)
		return
	}

	resp, err := profileSrv.ReadProfile(&id)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func updateProfile(c *gin.Context) {
	var updateProfileRequest *profiles.Profile
	err := c.Bind(&updateProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	if updateProfileRequest.UID == nil {
		updateProfileRequest.UID = c.GetString("UID")
	} else if updateProfileRequest.UID != c.GetString("UID") {
		c.JSON("402", gin.H{"error", "token UID does not match request UID"})
	}

	resp, err := profileSrv.UpdateProfile(updateProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

// func deleteProfile(c *gin.Context) {
// 	id := c.GetString("UID")

// 	resp, err := profileSrv.DeleteProfile(&id)
// 	if err != nil {
// 		c.JSON(501, err)
// 		return
// 	}

// 	c.JSON(200, &resp)
// }

func deleteJob(c *gin.Context) {

	req := profiles.DeleteJobRequest{}
	if req.UID == nil {
		req.UID = c.GetString("UID")
	} else if req.UID != c.GetString("UID") {
		c.JSON("402", gin.H{"error", "token UID does not match request UID"})
	}

	resp, err := profileSrv.DeleteJob(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteEdu(c *gin.Context) {

	req := profiles.DeleteEduRequest{}
	if req.UID == nil {
		req.UID = c.GetString("UID")
	} else if req.UID != c.GetString("UID") {
		c.JSON("402", gin.H{"error", "token UID does not match request UID"})
	}

	resp, err := profileSrv.DeleteEdu(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readAllProfiles(c *gin.Context) {
	resp, err := profileSrv.ReadAllProfiles()
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func findMentors(c *gin.Context) {
	uid := c.GetString("UID")
	resp, err := profileSrv.FindMentor(uid)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func requestMentor(c *gin.Context) {
	req := &profiles.MentorRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	resp, err := profileSrv.RequestMentor(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteMentor(c *gin.Context) {
	req := &profiles.MentorRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	resp, err := profileSrv.DeleteMentor(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)

}

func viewMentorRequests(c *gin.Context) {
	uid := c.GetString("UID")
	resp, err := profileSrv.FindMentor(uid)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func handleMentorRequest(c *gin.Context) {
	req := &profiles.MentorRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	resp, err := profileSrv.DeleteMentor(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteMentee(c *gin.Context) {
	req := &profiles.MentorRequest{}

	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	resp, err := profileSrv.DeleteMentee(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func getMessages(c *gin.Context) {
	uid := c.GetString("UID")
	resp, err := profileSrv.GetMessages(uid)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)

}

func sendMessage(c *gin.Conext) {
	uid := c.GetString("UID")

	req := &profiles.Message{}
	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}
	req.From = uid
	resp, err := profileSrv.SendMessage(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}
