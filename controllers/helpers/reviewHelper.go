package helpers

import (
	"fmt"
	"localArtisans/configs"
	"localArtisans/models/database"
	"localArtisans/models/outputs"
	"localArtisans/models/repositories"
	"localArtisans/models/requestsDTO"
	"localArtisans/models/responsesDTO"

	"github.com/google/uuid"
)

func GetAllReviews(GetAllReviewsRequestDTO requestsDTO.GetAllReviewsRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var reviews []database.Reviews

	if GetAllReviewsRequestDTO.Limit > 100 {
		output := outputs.BadRequestOutput{
			Code: 400,
			Message: "Bad Request: Limit can't more than 100",
		}
		return 400, output
	}

	offset := (GetAllReviewsRequestDTO.Page - 1) * GetAllReviewsRequestDTO.Limit
	order := fmt.Sprintf("%s %s", GetAllReviewsRequestDTO.OrderBy, GetAllReviewsRequestDTO.OrderType)
	err := db.Offset(offset).Limit(GetAllReviewsRequestDTO.Limit).Order(order).Find(&reviews).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(reviews) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Reviews not exist",
		}
		return 404, output
	}

	var totalData int64
	var totalPage int
	db.Model(&responsesDTO.ReviewsResponseDTO{}).Count(&totalData)
	if totalData%int64(GetAllReviewsRequestDTO.Limit) == 0 {
		totalPage = int(totalData / int64(GetAllReviewsRequestDTO.Limit))
	} else {
		totalPage = int(totalData / int64(GetAllReviewsRequestDTO.Limit)) + 1
	}

	output := outputs.GetAllReviewsOutput{}
	output.Page = GetAllReviewsRequestDTO.Page	
	output.Limit = GetAllReviewsRequestDTO.Limit
	output.OrderBy = GetAllReviewsRequestDTO.OrderBy
	output.OrderType = GetAllReviewsRequestDTO.OrderType
	output.Code = 200
	output.Message = "Success: Reviews Found"
	output.TotalData = int(totalData)
	output.TotalPage = totalPage

	for _, review := range reviews {
		output.Data = append(output.Data, responsesDTO.ReviewsResponseDTO{
			ID: review.ID,
			ProductID: review.ProductID,
			UserID: review.UserID,
			Comment: review.Comment,
			Rating: review.Rating,
			IsActive: review.IsActive,
			CreatedBy: review.CreatedBy,
			UpdatedBy: review.UpdatedBy,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		})
	}
	return 200, output
}

func GetReviewByID(id string) (int, interface{}){
	var review database.Reviews
	review, err := repositories.GetReviewByID(id)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if review.ID == uuid.Nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Review not exist",
		}
		return 404, output
	}

	output := outputs.GetReviewOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success",
		},
		Data: responsesDTO.ReviewsResponseDTO{
			ID: review.ID,
			ProductID: review.ProductID,
			UserID: review.UserID,
			Comment: review.Comment,
			Rating: review.Rating,
			IsActive: review.IsActive,
			CreatedBy: review.CreatedBy,
			UpdatedBy: review.UpdatedBy,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		},
	}

	return 200, output
}


func GetAllReviewsByProductID(productID string) (int, interface{}){
	var reviews []responsesDTO.ReviewsResponseDTO
	reviews, err := repositories.GetAllReviewsByProductID(productID)

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	if len(reviews) == 0 {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Reviews not exist",
		}
		return 404, output
	}
	output := outputs.GetAllReviewsByProductIDOutput{}
	output.Code = 200
	output.Message = "Success: Reviews Found"
	output.Data = reviews
	
	return 200, output
}

func CreateReview(CreateReviewRequestDTO requestsDTO.CreateReviewRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	review := database.Reviews{
		ProductID: CreateReviewRequestDTO.ProductID,
		UserID: CreateReviewRequestDTO.UserID,
		Comment: CreateReviewRequestDTO.Comment,
		Rating: CreateReviewRequestDTO.Rating,
		IsActive: true,
		CreatedBy: CreateReviewRequestDTO.CreatedBy,
	}
	err := db.Create(&review).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	var product database.Products

	err = db.Table("products").Where("id = ?", CreateReviewRequestDTO.ProductID).First(&product).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Product not exist",
		}
		return 404, output
	}

	// rating hanya bisa 1 - 5, gk bisa kasih 0 atau lebih dari 5
	if product.Rating == 0 {
		product.Rating = (product.Rating + CreateReviewRequestDTO.Rating)
	} else {
		product.Rating = (product.Rating + CreateReviewRequestDTO.Rating) / 2
	}

	err = db.Save(&product).Error

	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.CreateReviewOutput{
		BaseOutput: outputs.BaseOutput{
			Code: 200,
			Message: "Success: Review Created",
		},
		Data: responsesDTO.ReviewsResponseDTO{
			ID: review.ID,
			ProductID: review.ProductID,
			UserID: review.UserID,
			Comment: review.Comment,
			Rating: review.Rating,
			IsActive: review.IsActive,
			CreatedBy: review.CreatedBy,
			UpdatedBy: review.UpdatedBy,
			CreatedAt: review.CreatedAt,
			UpdatedAt: review.UpdatedAt,
		},
	}
	return 200, output
}

func DeleteReview(DeleteReviewRequestDTO requestsDTO.DeleteReviewRequestDTO) (int, interface{}) {
	db := configs.GetDB()
	var review database.Reviews
	err := db.Where("id = ?", DeleteReviewRequestDTO.ID).First(&review).Error

	if err != nil {
		output := outputs.NotFoundOutput{
			Code: 404,
			Message: "Not Found: Review not exist",
		}
		return 404, output
	}

	err = db.Delete(&review).Error
	if err != nil {
		output := outputs.InternalServerErrorOutput{
			Code: 500,
			Message: "Internal Server Error" + err.Error(),
		}
		return 500, output
	}

	output := outputs.DeleteReviewOutput{}
	output.Code = 200
	output.Message = "Success: Review deleted"
	output.Data = responsesDTO.ReviewsResponseDTO{
		ID: review.ID,
		ProductID: review.ProductID,
		UserID: review.UserID,
		Comment: review.Comment,
		Rating: review.Rating,
		IsActive: review.IsActive,
		CreatedBy: review.CreatedBy,
		UpdatedBy: review.UpdatedBy,
		CreatedAt: review.CreatedAt,
		UpdatedAt: review.UpdatedAt,
	}

	return 200, output
}