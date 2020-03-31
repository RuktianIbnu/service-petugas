package module

import (
	"net/http"
    "../structs"
    "context"
    "github.com/nbys/asyncwork/worker"
    
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
)

func (idb *InDB) InsertDataPerubahan(c *gin.Context){
	var (
        data structs.EditPermohonan
	)

	id_permohonan     	:= c.PostForm("ID_Permohonan")
    id_user_pemohon     := c.PostForm("id_user_pemohon")
    nama_server         := c.PostForm("nama_server")
    detail              := c.PostForm("detail")
    jenis_server        := c.PostForm("jenis_server")
    os                  := c.PostForm("os")
    ram                 := c.PostForm("ram")
    storage             := c.PostForm("storage")
    hostname            := c.PostForm("hostname")
    internet            := c.PostForm("internet")
    internet_status     := c.PostForm("internet_status")
    open_port           := c.PostForm("open_port")
    lokasi              := c.PostForm("lokasi")
    id_kontainment      := c.PostForm("id_kontainment")
    rak                 := c.PostForm("rak")

	data.ID_Permohonan		= id_permohonan
    data.ID_user_pemohon    = id_user_pemohon
    data.Status             = "APPROVE_PETUGAS"
    data.Nama_server        = nama_server
    data.Detail             = detail
    data.Jenis_server       = jenis_server
    data.Os                 = os
    data.Ram                = ram
    data.Storage            = storage
    data.Hostname           = hostname
    data.Internet           = internet
    data.Internet_status    = internet_status
    data.Open_port          = open_port
    data.Lokasi             = lokasi
    data.Id_kontainment     = id_kontainment
    data.Rak                = rak 

    valid := CekAuth(c)

    if valid == true {
        success := idb.DB.Create(&data).Error
        if success != nil {
            c.JSON(http.StatusInternalServerError, gin.H {
                "pesan": "gagal simpan data",
                "status": "error",
            })
            c.Abort()
        } else {
            c.JSON(http.StatusOK, gin.H {
                "status": "success",
                "pesan": "Data berhasil di simpan",
            })
        }
    } else {
        c.JSON(http.StatusOK, gin.H {
            "status": "error",
			"pesan": "not authorized",
        })
    }
    
}

func (idb *InDB) InsertReview(c *gin.Context){
	var (
        data structs.ReviewPetugas
        dataIP structs.Ip_Address
        dataManagement structs.Management
        dataSpesifikasi structs.Spesifikasi

        hasil gin.H
	)

    ID_permohonan     	    := c.PostForm("id_permohonan")
	ID_petugas     	        := c.PostForm("id_petugas")
    Status_review           := "NEW"
    data.ID_permohonan      = ID_permohonan
    data.ID_petugas         = ID_petugas
    data.Status_review      = Status_review

    Ip_address              := c.PostForm("ip_address")
    Ip_public               := c.PostForm("ip_public")
    Hostname                := c.PostForm("hostname")
    Netmask                 := c.PostForm("netmask")
    Geteway                 := c.PostForm("geteway")
    Dns                     := c.PostForm("dns")
    dataIP.Ip_address       = Ip_address
    dataIP.Ip_public        = Ip_public
    dataIP.Hostname         = Hostname
    dataIP.Netmask          = Netmask
    dataIP.Geteway          = Geteway
    dataIP.Dns              = Dns
    dataIP.ID_permohonan    = ID_permohonan

    Ip_address_management           := c.PostForm("ip_address_management")
    Username_management             := c.PostForm("username_management")
    Password_management             := c.PostForm("hostname_management")
    Netmask_management              := c.PostForm("netmask_management")
    Geteway_management              := c.PostForm("geteway_management")
    Dns_management                  := c.PostForm("dns_management")
    dataManagement.Ip_address       = Ip_address_management
    dataManagement.Username         = Username_management
    dataManagement.Password         = Password_management
    dataManagement.Netmask          = Netmask_management
    dataManagement.Geteway          = Geteway_management
    dataManagement.Dns              = Dns_management
    dataManagement.ID_permohonan    = ID_permohonan

    Jenis_server                    := c.PostForm("jenis_server")
    Os                              := c.PostForm("os")
    Ram                             := c.PostForm("ram")
    Storage                         := c.PostForm("storage")
    dataSpesifikasi.Jenis_server    = Jenis_server
    dataSpesifikasi.Os              = Os
    dataSpesifikasi.Ram             = Ram
    dataSpesifikasi.Storage         = Storage
    dataSpesifikasi.ID_permohonan   = ID_permohonan

    valid := CekAuth(c)

    if valid == true {
        task1 := func() interface{} {
            hasil := idb.ExeInsertReview(data, c)
            return hasil 
        }
        task2 := func() interface{} {
            hasil2 := idb.ExeInsertIP(dataIP, c)
            return hasil2
        }
        task3 := func() interface{} {
            hasil3 := idb.ExeInsertManagement(dataManagement, c)
            return hasil3
        }
        task4 := func() interface{} {
            hasil4 := idb.ExeInsertSpesifikasi(dataSpesifikasi, c)
            return hasil4
        }

        tasks := []worker.TaskFunction{task1, task2, task3, task4}

        ctx, cancel := context.WithCancel(context.Background())
        defer cancel()
        
        resultChannel := worker.PerformTasks(ctx, tasks)

        for result := range resultChannel {
            switch {
            case result == "error":
                hasil = gin.H {
                    "status": "warning!",
                        "pesan": "Silahkan lengkapi lembar isian",
                }
                cancel()
                return
            case result == "success":
                hasil = gin.H {
                        "pesan": "Berhasil",
                        "status": result,
                }
            default:
                hasil = gin.H {
                    "pesan": "gagal",
                    "status": "error",
                    "result": result,
                }
            }
            c.JSON(http.StatusOK, hasil)
        }
    } else {
        c.JSON(http.StatusOK, gin.H {
            "status": "error",
			"pesan": "not authorized",
        })
    }
}