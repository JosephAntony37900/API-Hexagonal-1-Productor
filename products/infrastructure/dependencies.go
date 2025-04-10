package infrastructure

import (
	"database/sql"
	"github.com/gin-gonic/gin"
    productApp "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/application"
    productController "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/controllers"
    productRepo "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/repository"
    productRoutes "github.com/JosephAntony37900/API-Hexagonal-1-Productor/products/infrastructure/routes"
)

func InitProductDependencies(engine *gin.Engine, db *sql.DB) {


	productRepository := productRepo.NewProductRepoMySQL(db)

	createProduct := productApp.NewCreateProduct(productRepository)
    getProducts := productApp.NewGetProducts(productRepository)
    updateProduct := productApp.NewUpdateProduct(productRepository)
    deleteProduct := productApp.NewDeleteProduct(productRepository)
    getProductsByMinPrice := productApp.NewGetProductsByMinPrice(productRepository)
	getProductById := productApp.NewGetProductByID(productRepository)

	createProductController := productController.NewCreateProductController(createProduct)
    getProductsController := productController.NewGetProductsController(getProducts)
    updateProductController := productController.NewUpdateProductController(updateProduct)
    deleteProductController := productController.NewDeleteProductController(deleteProduct)
    getProductsByMinPriceController := productController.NewGetProductsByMinPriceController(getProductsByMinPrice)
	getProductByIdController := productController.NewGetProductByIDController(getProductById)

	productRoutes.SetupProductRoutes(engine, createProductController, getProductsController, updateProductController, deleteProductController, getProductsByMinPriceController, getProductByIdController)


}