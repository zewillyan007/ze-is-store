package db

import "fmt"

func DeleteProduct(id string) {
	db := ConectionMySql()

	deleteTheProduct, err := db.Prepare("delete from produtos where id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err2 := deleteTheProduct.Exec(id)

	if err2 != nil {
		fmt.Printf("%#v", err2)
	}
	defer db.Close()
}
