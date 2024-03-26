package main

func init(){
	ConnectToDB()

}

func main(){

	db.AutoMigrate(&user{})
}