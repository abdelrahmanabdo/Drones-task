package drone

import (
	"github.com/gin-gonic/gin"
)

func createDrone(c *gin.Context) error {
	u := &drone{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	drones[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

func getDrone(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, drones[id])
}

func getDrones(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, drones[id])
}

func updateDrone(c *gin.Context) error {
	u := new(drone)
	if err := c.Bind(u); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	drones[id].Name = u.Name
	return c.JSON(http.StatusOK, drones[id])
}

func deleteDrone(c *gin.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(drones, id)
	return c.NoContent(http.StatusNoContent)
}