package controlfloor_auth

import (
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/alexedwards/scs/v2"
    uj "github.com/nanoscopic/ujsonin/v2/mod"
)

type AuthHandler interface {
    UserAuth( c *gin.Context ) bool
    UserLogin( c *gin.Context ) bool
}

type SessionManager interface {
    GetSCSSessionManager() *scs.SessionManager
    GetSession( c *gin.Context ) context.Context
    WriteSession( c *gin.Context )
}

type authUser struct {
    userName string
    password string
}

type demoAH struct {
    sessionManager SessionManager
    testParam      string
    testUser       string
    users          []authUser
}

func NewAuthHandler( confRoot uj.JNode, sessionManager SessionManager ) AuthHandler {
    self := &demoAH{
        sessionManager: sessionManager,
        testParam: "",
        testUser: "test",
        users: []authUser{},
    }
       
    return self
}

func (self *demoAH) UserAuth( c *gin.Context ) bool {
    fmt.Printf("uauth\n")
    
    // Force authentication
    sm := self.sessionManager
    s := sm.GetSession(c)
    scsSM := sm.GetSCSSessionManager()
    scsSM.Put( s, "user", "any" )
    sm.WriteSession( c )
        
    c.Next()
    return true
}

func (self *demoAH) UserLogin( c *gin.Context ) bool {
    s := self.sessionManager.GetSession( c )
    scsSM := self.sessionManager.GetSCSSessionManager()
    
    // Accepting any given username blindly for login
    user := c.PostForm("user")
    
    fmt.Printf( "login ok; user=%s\n", user )
    
    scsSM.Put( s, "user", user )
    self.sessionManager.WriteSession( c )
            
    return true
}
