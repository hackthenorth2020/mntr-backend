package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hackthenorth2020/mntr-backend/profiles"
)

func createProfile(c *gin.Context) {
	var createProfileRequest *profiles.Profile
	err := c.Bind(&createProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	if createProfileRequest.UID == "" {
		createProfileRequest.UID = c.GetString("UID")
	} else if createProfileRequest.UID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
	}

	resp, err := profileSrv.CreateProfile(createProfileRequest)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func readProfile(c *gin.Context) {
	uid := c.Param("id")

	resp, err := profileSrv.ReadProfile(uid)
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

	if updateProfileRequest.UID == "" {
		updateProfileRequest.UID = c.GetString("UID")
	} else if updateProfileRequest.UID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
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

	req := &profiles.DeleteJobRequest{}
	if req.UID == "" {
		req.UID = c.GetString("UID")
	} else if req.UID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
	}

	resp, err := profileSrv.DeleteJob(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func deleteEdu(c *gin.Context) {

	req := &profiles.DeleteEduRequest{}
	if req.UID == "" {
		req.UID = c.GetString("UID")
	} else if req.UID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
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
		log.Printf("[ERROR] [FIND MENTORS] %v", err)
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

	if req.MenteeUID == "" {
		req.MenteeUID = c.GetString("UID")
	}
	// else if req.MenteeUID != c.GetString("UID") {
	// 	c.JSON(402, "token UID does not match request UID")
	// 	return
	// }

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

	if req.MentorUID == "" {
		req.MentorUID = c.GetString("UID")
	}
	if req.MenteeUID == "" {
		req.MenteeUID = c.GetString("UID")
	}

	if req.MentorUID == req.MenteeUID {
		c.JSON(402, "why is mentorUID = menteeUID??? Maybe both are empty?")
		return
	} else if req.MentorUID != c.GetString("UID") && req.MenteeUID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
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
	resp, err := profileSrv.ViewMentorRequests(uid)
	if err != nil {
		log.Printf("[ERROR] [VIEW MENTOR REQS] %v", err)
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

func handleMentorRequest(c *gin.Context) {
	req := &profiles.MentorResponse{}

	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	if req.MentorUID == "" {
		req.MentorUID = c.GetString("UID")
	} else if req.MentorUID != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID | "+req.MentorUID+" | "+c.GetString("UID"))
	}

	resp, err := profileSrv.HandleMentorRequest(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}

// func deleteMentee(c *gin.Context) {
// 	req := &profiles.MentorRequest{}

// 	if err := c.Bind(&req); err != nil {
// 		c.JSON(501, err)
// 		return
// 	}

// 	resp, err := profileSrv.DeleteMentee(req)
// 	if err != nil {
// 		c.JSON(501, err)
// 		return
// 	}

// 	c.JSON(200, &resp)
// }

func getMessages(c *gin.Context) {
	req := &profiles.GetMessageRequest{}
	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}

	if req.From == "" {
		req.From = c.GetString("UID")
	} else if req.From != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
	}
	resp, err := profileSrv.GetMessages(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)

}

func sendMessage(c *gin.Context) {
	req := &profiles.Message{}
	if err := c.Bind(&req); err != nil {
		c.JSON(501, err)
		return
	}
	if req.From == "" {
		req.From = c.GetString("UID")
	} else if req.From != c.GetString("UID") {
		c.JSON(402, "token UID does not match request UID")
	}
	resp, err := profileSrv.SendMessage(req)
	if err != nil {
		c.JSON(501, err)
		return
	}

	c.JSON(200, &resp)
}
