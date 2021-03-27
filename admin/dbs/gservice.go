package dbs

import (
	"fmt"
)

// g 服务表结构
type GService struct {
    Id int64 `db:"id"`
    Name string  `db:"name"`
    Status int `db:"status"`
    Host string  `db:"host"`
    Port string  `db:"port"`
    DockerImage string  `db:"docker_image"`
    ContainerId string  `db:"container_id"`
    serviceType string  `db:"service_type"`
}



// 查询数据，指定字段名
func GetGService(id string) {

    service := new(GService)
    row := MysqlDb.QueryRow("select id, name, status, host, port, docker_image, container_id,service_type from g_service where id=?",id)
    if err :=row.Scan(&service.Id,&service.Name,&service.Status, &service.Host,&service.Port,&service.DockerImage,&service.ContainerId,&service.serviceType); err != nil{
        fmt.Println("scan failed, err: ",err)
        return
    }

    fmt.Println(service.Id,service.Name,service.Status,service.Host,service.Port, service.DockerImage, service.ContainerId);
}


// 查询数据，取所有字段
func GetGServices() {

    // 通过切片存储
    servcies := make([]GService, 0)
    rows, _:= MysqlDb.Query("SELECT * FROM `g_service` limit ?",100)
    // 遍历
    var gservice GService
    for  rows.Next(){
        rows.Scan(&gservice.Id, &gservice.Name, &gservice.Status, &gservice.Host, &gservice.Port, &gservice.DockerImage, &gservice.ContainerId, &gservice.serviceType)
        servcies=append(servcies,gservice)
    }

    fmt.Println(servcies)

}


// 插入数据
func CreateGservice(name string, status int, host string, port string, docker_image string, container_id string, service_type string) (string, error) {

    ret,err := MysqlDb.Exec("insert INTO g_service(name, status, host, port, docker_image, container_id,service_type) values(?,?,?,?,?,?,?)", 
                                                name, status, host, port, docker_image, container_id, service_type)

    if err != nil {
        fmt.Println("machine db exec fail", err)
        return "", err
    }

    //插入数据的主键id
    lastInsertID,_ := ret.LastInsertId()
    fmt.Println("LastInsertID:",lastInsertID)

    //影响行数
    rowsaffected,_ := ret.RowsAffected()
    fmt.Println("RowsAffected:",rowsaffected)

    return fmt.Sprint(lastInsertID),  nil

}

func DeleteService(container_id string) {

    ret,err := MysqlDb.Exec("delete from g_service where container_id=?", container_id)
    if err!=nil {
        fmt.Println("del service err", err)
    }    
    del_nums,_ := ret.RowsAffected()
    fmt.Println("del:",del_nums)
}






