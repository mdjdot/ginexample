package main

import "github.com/gin-gonic/gin"

// StructA .
type StructA struct {
	FieldA string `form:"field_a"`
}

// StructB .
type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

// StructC .
type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

// StructD .
type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

// GetDataB .
func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

// GetDataC .
func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

// GetDataD .
func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	r.Run()
}
