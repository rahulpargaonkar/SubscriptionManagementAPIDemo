# SubscriptionManagementAPIDemo

- A REST API solution for Subscription Management using GOlang, Gin & Gorm
Sample APIS:
1. Category:
  a. Add Category 
    - /api/category/addCategory
  b. Update Category
    - /api/updateCategory/:categoryName
  c. 
  
2. Subcategory
  a. Add SubCategory
     - /api/subcategory/addSubcategories
  b. Update SubCategory
     - /api/subcategory/updateSubcategory/:categoryName/:subcategoryName

3. Product
  a. Add Product 
     - /api/product/addProduct/:categoryName/:subcategoryName
  b. Add Weeklyprices for Product
     - /api/product/addProductsWeeklyPrices/:productName
    
4. User
  a. Add New User
     - /api/user/addUser
  b. Get All Active Subscription of User
     - /api/user/GetAllActiveSubscriptions/:email
   
5. Subscription
  a. Get Subscription Amount of product for given duration
     - /getSubScriptionAmount/:productName/:startDate/:subscriptionType
  b. Add Subscription of specific product for Perticular User
     - /AddSubscription/:email/:productName/:subscriptionType/:startDate
  
