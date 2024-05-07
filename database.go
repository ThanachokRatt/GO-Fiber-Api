package main

func createProduct(product *Product) error{
	//INSERT INTO public.products(id, names, price, supplier_id)VALUES (?, ?, ?, ?);

     _, err := db.Exec("INSERT INTO public.products(names, price)VALUES ($1, $2);",
	 product.Name,
	 product.Price,

    )
    return err

}

func getProduct(id int) (Product, error){
	var p Product
	row := db.QueryRow("SELECT id,names,price FROM products WHERE id=$1;",id,)
	err := row.Scan(&p.ID,&p.Name,&p.Price)

	if err != nil {
		return Product{},err
	}
	return p,nil

}

func getProducts() ([]Product,error){
	rows,err := db.Query("SELECT id,names,price from products")

	if err != nil {
		return nil,err
	}

	var products []Product

	for rows.Next(){
		var p Product
		err := rows.Scan(&p.ID,&p.Name,&p.Price)
		if err != nil {
			return nil,err
		}
		products = append(products, p)
	}

	if err = rows.Err(); err != nil {
		return nil,err
	}
	return products,nil
}



func updateProduct(id int,product *Product) (Product,error){
	var p Product
	//UPDATE public.productsSET id=?, names=?, price=?, supplier_id=?WHERE <condition>;

     row := db.QueryRow("UPDATE public.products SET names=$1, price=$2 WHERE id=$3 RETURNING id,names,price;",
	 product.Name,
	 product.Price,
	 id,

    )
    err := row.Scan(&p.ID,&p.Name,&p.Price)

	if err != nil {
		return Product{},err
	}

    return p,err

    
}

func deleteProduct(id int) error {
     _,err := db.Exec("DELETE FROM public.products WHERE id=$1;",id,)
	return err

}

