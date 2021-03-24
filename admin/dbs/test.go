package dbs

import (
	"fmt"
	"log"
)

// 用户表结构体
type User struct {
    Id int64 `db:"id"`
    Name string  `db:"name"`
    Age int `db:"age"`
}

// 查询数据，指定字段名
func StructQueryField() {

    user := new(User)
    row := MysqlDb.QueryRow("select id, name, age from users where id=?",1)
    if err :=row.Scan(&user.Id,&user.Name,&user.Age); err != nil{
        fmt.Printf("scan failed, err:%v",err)
        return
    }
    fmt.Println(user.Id,user.Name,user.Age)
}

// 查询数据，取所有字段
func StructQueryAllField() {

    // 通过切片存储
    users := make([]User, 0)
    rows, _:= MysqlDb.Query("SELECT * FROM `users` limit ?",100)
    // 遍历
    var user User
    for  rows.Next(){
        rows.Scan(&user.Id, &user.Name, &user.Age)
        users=append(users,user)
    }

    fmt.Println(users)

}

// 插入数据
func StructInsert() {

    ret,err := MysqlDb.Exec("insert INTO users(name,age) values(?,?)","小红",23)

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

// 更新数据
func StructUpdate() {

    ret,_ := MysqlDb.Exec("UPDATE users set age=? where id=?","100",1)
    upd_nums,_ := ret.RowsAffected()

    fmt.Println("RowsAffected:",upd_nums)
}

// 删除数据
func StructDel() {

    ret,_ := MysqlDb.Exec("delete from users where id=?",1)
    del_nums,_ := ret.RowsAffected()

    fmt.Println("RowsAffected:",del_nums)
}


// 事务处理,结合预处理
func StructTx() {

    //事务处理
    tx, _ := MysqlDb.Begin();


    // 新增
    userAddPre, _ := MysqlDb.Prepare("insert into users(name, age) values(?, ?)");
    addRet, _ := userAddPre.Exec("zhaoliu", 15);
    ins_nums, _ := addRet.RowsAffected();


    // 更新
    userUpdatePre1, _ := tx.Exec("update users set name = 'zhansan'  where name=?", "张三");
    upd_nums1, _ := userUpdatePre1.RowsAffected();
    userUpdatePre2, _ := tx.Exec("update users set name = 'lisi'  where name=?", "李四");
    upd_nums2, _ := userUpdatePre2.RowsAffected();

    fmt.Println(ins_nums);
    fmt.Println(upd_nums1);
    fmt.Println(upd_nums2);

    if ins_nums > 0 && upd_nums1 > 0 && upd_nums2 > 0 {
        tx.Commit();
    }else{
        tx.Rollback();
    }

}

// 查询数据，指定字段名,不采用结构体
func  RawQueryField() {

    rows, _ := MysqlDb.Query("select id,name from users");
    if rows == nil {
        log.Fatal("mysql exe fail")
        return
    }
    id := 0;
    name := "";
    for rows.Next() {
        rows.Scan(&id, &name);
        fmt.Println(id, name);
    }
}

// 查询数据,取所有字段,不采用结构体
func RawQueryAllField() {

    //查询数据，取所有字段
    rows2, _ := MysqlDb.Query("select * from users");

    //返回所有列
    cols, _ := rows2.Columns();

    //这里表示一行所有列的值，用[]byte表示
    vals := make([][]byte, len(cols));

    //这里表示一行填充数据
    scans := make([]interface{}, len(cols));
    //这里scans引用vals，把数据填充到[]byte里
    for k, _ := range vals {
        scans[k] = &vals[k];
    }

    i := 0;
    result := make(map[int]map[string]string);
    for rows2.Next() {
        //填充数据
        rows2.Scan(scans...);
        //每行数据
        row := make(map[string]string);
        //把vals中的数据复制到row中
        for k, v := range vals {
            key := cols[k];
            //这里把[]byte数据转成string
            row[key] = string(v);
        }
        //放入结果集
        result[i] = row;
        i++;
    }
    fmt.Println(result);
}