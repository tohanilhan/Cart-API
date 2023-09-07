
# Cart API

---

*Simple Shopping Cart API that has some business rules such as discounts. It comes with a PostgresSQL to store data and has necessary files for containerization and **Makefile** for deployment.*

---
# About the project

*The functions of this service are as follows;*

1. List Products
   - Users should be able to list all products.
2. Add To Cart
   - Users can add their products to the basket and the total of the basket changes accordingly.
3. Show Cart
   - Users can list the products they have added to their cart and total price and VAT of the cart.
4. Delete Cart Item
   - Users can remove the products added from the cart. Notice removing an item may change discount.
5. Complete Order
   - Users can create an order with the products they add to their cart. There is no need to implement any payment mechanism. You can assume that when an order is completed it is fully paid.

## Some business rules

1. Products always have price and VAT (Value Added Tax, or KDV). VAT might be different for different products. Typical VAT percentage is %1, %8 and %18. So use these values for your products.

2. There might be discount in following situations:
   1.  Every fourth order whose total is more than given amount may have discount
depending on products. Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
   2. If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
   3. c. If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
   4. Only one discount can be applied at a time. Only the highest discount should be applied.

---

### Prerequisites and Installation

- #### **Docker:** 
  You can install Docker Desktop by following the instructions on the [Docker Desktop website](https://desktop.docker.com/).

  If you are using **MacOS**, you can install Docker by following commands:
    - *`brew install docker`*
    - *`docker run hello-world`*
  
 ---
- #### **Go:**
  
  If you haven't done already, You need to install Go by following the instructions on the [Go website](https://golang.org/doc/install).

---

- #### **Project:**
  
  You have to clone the project from the [Github repository](https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project) and run the below script to build the project
   
   *`git clone https://github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project.git`*

---

#### **Run the Project:**

You have to make sure that you are in the **same directory as the project**. Then you can run the below script:

####  **Makefile** 
- ##### Commands
  - **deploy-all** : Deploy all services
  - **deploy-postgresql** : Deploy postgresql service
  - **deploy-api** : Deploy cart-api service
  - **build-all**: Build all images
  - **build-postgres** : Build postgresql image
  - **build-api** : Build cart-api image
  - **clean-all** : Clean all images and containers
  - **clean-images** : Clean all images
  - **clean-containers** : Clean all containers
  - **clean-postgres** : Clean postgresql image and container
  - **clean-api** : Clean cart-api image and container
  - **clean-volumes** : Clean all volumes
  - **clean-networks** : Clean all networks
  - **down-all** : Down all services
  - **down-postgres** : Down postgresql service
  - **up-all** : Up all services
  - **down-api** : Down cart-api service

- ##### Usage
    ```make deploy-all``` 


- Or you can run the project directly by running the below script from the cart-api directory inside the project:

    ```go run main.go```

**Note:** If you are using **VS Code** as your editor, you can simply run this project with debug mode by pressing **F5**.

---

### Endpoints and their functionality

**1. Get all products**
   
**Endpoint:** `http://127.0.0.1:8080/api/v1/products-api/products//products/`

**Functionality:** *Get all products from the database.*

**Method:** GET

**2. Add product to basket**
   
**Endpoint:** `http://127.0.0.1:8080/api/v1/products-api/products/products/`

**Functionality:** *Add product to the cart by sending the product id.*

**Method:** POST

**3. Show basket**
   
**Endpoint:** `http://127.0.0.1:8080/api/v1/products-api/products/products/basket/`

**Functionality:** *Show the products in the basket.*

**Method:** POST

**4. Remove product from basket**
   
**Endpoint:** `http://127.0.0.1:8080/api/v1/products-api/products/products/basket/`

**Functionality:** *Delete product from the cart by sending the product id.*

**Method:** DELETE

**5. Complete order**
   
**Endpoint:**`http://127.0.0.1:8080/api/v1/products-api/products/products/order`

**Functionality:** *Complete the order. Under some circumstances, discount can be applied. If any discount can be applied, apply the discount and finalize the order.*

**Method:** POST

**6. Ping**

**Endpoint:**`http://127.0.0.1:8080/api/v1/ping`

**Functionality:** *Check if the service is up.*

**Method:** GET

---
