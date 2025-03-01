package main

import (
	"project-its/controllers"
	"project-its/initializers"
	// "project-its/middleware"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {

	initializers.LoadEnvVariables()
	initializers.ConnectToDB()

}

func main() {

	r := gin.Default()

	// Enable CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8000"
		},
		MaxAge: 12 * time.Hour,
	}))

	// Middleware untuk autentikasi token
	// authMiddleware := middleware.TokenAuthMiddleware()

	// Route yang tidak memerlukan autentikasi
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Terapkan middleware autentikasi ke semua route selanjutnya
	// r.Use(authMiddleware)
	// Routes for User
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/updateAll", controllers.UpdateAllSheets)
	r.GET("/exportAll", controllers.ExportAllSheets)

	// Setup session store
	store = cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	//logout must be after middleware
	r.POST("/logout", controllers.Logout)

	// Routes for SAG
	r.GET("/sag", controllers.SagIndex)
	r.POST("/sag", controllers.CreateSag)
	r.GET("/sag/:id", controllers.SagShow)
	r.PUT("/sag/:id", controllers.PostsUpdate)
	r.DELETE("/sag/:id", controllers.PostsDelete)
	r.GET("/exportSag", controllers.CreateExcelSag)
	r.GET("/updateSag", controllers.UpdateSheetSAG)
	r.POST("/uploadSag", controllers.ImportExcelSag)

	// Routes for Memo
	r.GET("/memos", controllers.MemoIndex)
	r.POST("/memos", controllers.MemoCreate)
	r.GET("/memos/:id", controllers.MemoShow)
	r.PUT("/memos/:id", controllers.MemoUpdate)
	r.DELETE("/memos/:id", controllers.MemoDelete)
	r.GET("/exportMemo", controllers.CreateExcelMemo)
	r.GET("/updateMemo", controllers.UpdateSheetMemo)
	r.POST("/uploadMemo", controllers.ImportExcelMemo)

	// Routes for ISO
	r.GET("/iso", controllers.IsoIndex)
	r.POST("/iso", controllers.IsoCreate)
	r.GET("/iso/:id", controllers.IsoShow)
	r.PUT("/iso/:id", controllers.IsoUpdate)
	r.DELETE("/iso/:id", controllers.IsoDelete)
	r.GET("/exportIso", controllers.CreateExcelIso)
	r.GET("/updateIso", controllers.UpdateSheetIso)
	r.POST("/uploadIso", controllers.ImportExcelIso)

	// Routes for Surat
	r.POST("/surat", controllers.SuratCreate)
	r.PUT("/surat/:id", controllers.SuratUpdate)
	r.GET("/surat", controllers.SuratIndex)
	r.DELETE("/surat/:id", controllers.SuratDelete)
	r.GET("/surat/:id", controllers.SuratShow)
	r.GET("/exportSurat", controllers.CreateExcelSurat)
	r.GET("/updateSurat", controllers.UpdateSheetSurat)
	r.POST("/uploadSurat", controllers.ImportExcelSurat)

	//BeritaAcara routes
	r.POST("/beritaAcara", controllers.BeritaAcaraCreate)
	r.PUT("/beritaAcara/:id", controllers.BeritaAcaraUpdate)
	r.GET("/beritaAcara", controllers.BeritaAcaraIndex)
	r.DELETE("/beritaAcara/:id", controllers.BeritaAcaraDelete)
	r.GET("/beritaAcara/:id", controllers.BeritaAcaraShow)
	r.GET("/exportBerita", controllers.CreateExcelBerita)
	r.GET("/updateBerita", controllers.UpdateSheetBerita)
	r.POST("/uploadBerita", controllers.ImportExcelBerita)

	//Sk routes
	r.POST("/sk", controllers.SkCreate)
	r.PUT("/sk/:id", controllers.SkUpdate)
	r.GET("/sk", controllers.SkIndex)
	r.DELETE("/sk/:id", controllers.SkDelete)
	r.GET("/sk/:id", controllers.SkShow)
	r.GET("/exportSk", controllers.CreateExcelSk)
	r.GET("/updateSk", controllers.UpdateSheetSk)
	r.POST("/uploadSk", controllers.ImportExcelSk)

	//Project routes
	r.POST("/Project", controllers.ProjectCreate)
	r.PUT("/Project/:id", controllers.ProjectUpdate)
	r.GET("/Project", controllers.ProjectIndex)
	r.DELETE("/Project/:id", controllers.ProjectDelete)
	r.GET("/exportProject", controllers.CreateExcelProject)
	r.GET("/updateProject", controllers.UpdateSheetProject)
	r.POST("/uploadProject", controllers.ImportExcelProject)

	// Ruang Rapat routes
	r.GET("/ruang-rapat", controllers.GetEvents)
	r.POST("/ruang-rapat", controllers.CreateEvent)
	r.DELETE("/ruang-rapat/:id", controllers.DeleteEvent)
	r.GET("/notifications", controllers.GetNotifications)
	r.PUT("/notifications/:id/read", controllers.MarkNotificationRead)
	r.DELETE("/notifications/:id", controllers.DeleteNotification)

	// Jadwal Cuti routes
	r.GET("/jadwal-cuti", controllers.GetEventsCuti)
	r.POST("/jadwal-cuti", controllers.CreateEventCuti)
	r.DELETE("/jadwal-cuti/:id", controllers.DeleteEventCuti)

	//Perdin routes
	r.POST("/Perdin", controllers.PerdinCreate)
	r.PUT("/Perdin/:id", controllers.PerdinUpdate)
	r.GET("/Perdin", controllers.PerdinIndex)
	r.DELETE("/Perdin/:id", controllers.PerdinDelete)
	r.GET("/Perdin/:id", controllers.PerdinShow)
	r.GET("/exportPerdin", controllers.CreateExcelPerdin)
	r.GET("/updatePerdin", controllers.UpdateSheetPerdin)
	r.POST("/uploadPerdin", controllers.ImportExcelPerdin)

	//Surat  Masuk routes
	r.POST("/SuratMasuk", controllers.SuratMasukCreate)
	r.PUT("/SuratMasuk/:id", controllers.SuratMasukUpdate)
	r.GET("/SuratMasuk", controllers.SuratMasukIndex)
	r.DELETE("/SuratMasuk/:id", controllers.SuratMasukDelete)
	r.GET("/SuratMasuk/:id", controllers.SuratMasukShow)
	r.GET("/exportSuratMasuk", controllers.CreateExcelSuratMasuk)
	r.GET("/updateSuratMasuk", controllers.UpdateSheetSuratMasuk)
	r.POST("/uploadSuratMasuk", controllers.ImportExcelSuratMasuk)

	//Surat  Keluar routes
	r.POST("/SuratKeluar", controllers.SuratKeluarCreate)
	r.PUT("/SuratKeluar/:id", controllers.SuratKeluarUpdate)
	r.GET("/SuratKeluar", controllers.SuratKeluarIndex)
	r.DELETE("/SuratKeluar/:id", controllers.SuratKeluarDelete)
	r.GET("/SuratKeluar/:id", controllers.SuratKeluarShow)
	r.GET("/exportSuratKeluar", controllers.CreateExcelSuratKeluar)
	r.GET("/updateSuratKeluar", controllers.UpdateSheetSuratKeluar)
	r.POST("/uploadSuratKeluar", controllers.ImportExcelSuratKeluar)

	r.Run()
}
