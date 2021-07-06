package controllers
 
import (
    “goblog/app/models”
 

    “golang.org/x/crypto/bcrypt”
    “github.com/revel/revel”
)
 
type App struct {
    GormController
}
 
func (c App) Login() revel.Result {
    return c.Render()
}
 
func (c App) CreateSession(username, password string) revel.Result {
    var user models.User
 

    // ➊ username으로 사용자 조회
    c.Txn.Where(&models.User{Username: username}).First(&user)
 

    // ➋ bcrypt 패키지의 CompareHashAndPassword 함수로 패스워드 비교
    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
 

    // ➌ 패스워드가 일치하면 세션 생성후 포스트 목록 화면으로 이동
    if err == nil {
        authKey := revel.Sign(user.Username)
        c.Session[“authKey”] = authKey
        c.Session[“username”] = user.Username
        c.Flash.Success(“Welcome, “ + user.Name)
        return c.Redirect(Post.Index)
    }
 

    // ➍ 세션 정보를 모두 제거하고 홈으로 이동
    for k := range c.Session {
        delete(c.Session, k)
    }
    c.Flash.Out[“username”] = username
    c.Flash.Error(“Login failed”)
    return c.Redirect(Home.Index)
}
 
func (c App) DestroySession() revel.Result {
    // clear session
    for k := range c.Session {
        delete(c.Session, k)
    }
    return c.Redirect(Home.Index)
}
