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
func GetGServices() {

    service := new(GService)
    row := MysqlDb.QueryRow("select id, name, status, host, port, docker_image, container_id,service_type from g_service where id=?",1)
    if err :=row.Scan(&service.Id,&service.Name,&service.Status, &service.Host,&service.Port,&service.DockerImage,&service.ContainerId,&service.serviceType); err != nil{
        fmt.Printf("scan failed, err:%v",err)
        return
    }
    fmt.Println(service.Id,service.Name,service.Status,service.Host,service.Port, service.DockerImage, service.ContainerId);
}


// 插入数据
func CreateGservice(name string, status int, host string, port string, docker_image string, container_id string, service_type string) {

    ret,err := MysqlDb.Exec("insert INTO g_service(name, status, host, port, docker_image, container_id,service_type) values(?,?,?,?,?,?,?)", 
                                                name, status, host, port, docker_image, container_id, service_type)

    if err != nil {
        fmt.Println("machine db exec fail", err)
        return
    }

    //插入数据的主键id
    lastInsertID,_ := ret.LastInsertId()
    fmt.Println("LastInsertID:",lastInsertID)

    //影响行数
    rowsaffected,_ := ret.RowsAffected()
    fmt.Println("RowsAffected:",rowsaffected)

}






