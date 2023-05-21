package courses

import (
	"net/http"
	"pator-api/models"
	"pator-api/utils/handlers"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	course := models.Course{}

	courses, err := course.FindAllCourses()

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
	}

	handlers.HandleSuccess(c, http.StatusOK, courses)
}

func Store(c *gin.Context) {
	var input CreateCourseValidation

	if err := c.ShouldBindJSON(&input); err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	course := models.Course{
		CourseName:  input.CourseName,
		CourseProdi: input.CourseProdi,
	}

	_, err := course.SaveCourse()

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	handlers.HandleSuccess(c, http.StatusCreated, nil)
}

func Show(c *gin.Context) {
	// get course id from url
	id := c.Param("id")

	course_id, err := handlers.StringToUint(id)

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}

	courseReceived, err := course.FindCourseByID(course_id)

	if err != nil {
		handlers.HandleError(c, http.StatusNotFound, err)
		return
	}

	handlers.HandleSuccess(c, http.StatusOK, courseReceived)
}

func Update(c *gin.Context) {
	// get course id from url
	id := c.Param("id")

	course_id, err := handlers.StringToUint(id)

	if err != nil {
		handlers.HandleError(c, http.StatusBadRequest, err)
		return
	}

	course := models.Course{}

	courseReceived, err := course.FindCourseByID(course_id)

	if err != nil {
		

}
