package query

var CATEGORY_SQL =
` 
  SELECT * FROM categories OFFSET $1 LIMIT $2
`
var PRODUCT_SQL =
`
  SELECT 
    product_id,
    name,
	  created_at
  FROM products
  WHERE category_id=$1

`