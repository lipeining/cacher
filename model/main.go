package model

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type MySQLTableType interface {
	CacheKey() string
	TableName() string
	MGet(opts ...DBOption) ([]MySQLTableType, error)
}

type DBOption func(*gorm.DB) *gorm.DB

type A struct {
	gorm.Model
	Name string
}

func (A) CacheKey() string {
	return "a"
}

func (A) TableName() string {
	return "a"
}

func (A) MGet(opts ...DBOption) ([]A, error) {
	db := &gorm.DB{}
	for _, opt := range opts {
		db = opt(db)
	}

	ret := make([]A, 0)

	err := db.Debug().Model(&A{}).Find(&ret).Error
	if err != nil {
		return nil, errors.Wrapf(err, "mget records from db fail")
	}

	return ret, nil
}

type BaseModel[T any] struct {
	DB *gorm.DB
}

func NewBaseModel[T any](db *gorm.DB) *BaseModel[T] {
	return &BaseModel[T]{DB: db}
}

func (b *BaseModel[T]) FindAll(where interface{}) ([]T, error) {
	var results []T
	err := b.DB.Where(where).Find(&results).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find all records fail")
	}
	return results, nil
}

func (b *BaseModel[T]) FindOne(where interface{}) (*T, error) {
	var result T
	err := b.DB.Where(where).First(&result).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find one record fail")
	}
	return &result, nil
}

func (b *BaseModel[T]) FindOrCreate(where interface{}, values interface{}) (*T, error) {
	var result T
	err := b.DB.Where(where).Attrs(values).FirstOrCreate(&result).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find or create record fail")
	}
	return &result, nil
}

type User struct {
	BaseModel[User]
	Username string
	Email    string
}

// 示例：创建一个新的用户
func CreateUser(db *gorm.DB, username, email string) (*User, error) {
	user := &User{
		BaseModel: *NewBaseModel[User](db), // 初始化 BaseModel
		Username:  username,
		Email:     email,
	}
	err := db.Create(user).Error
	if err != nil {
		return nil, errors.Wrapf(err, "create user fail")
	}
	return user, nil
}

// 示例：查找用户
func ExampleUsage(db *gorm.DB) {
	// 创建用户
	user, err := CreateUser(db, "john_doe", "john@example.com")
	if err != nil {
		// 处理错误
		return
	}

	// 查找用户
	var foundUser *User
	foundUser, err = user.FindOne(map[string]interface{}{"username": "john_doe"})
	if err != nil {
		// 处理错误
		return
	}

	// 输出找到的用户信息
	fmt.Printf("Found User: %s, Email: %s\n", foundUser.Username, foundUser.Email)
}

// IModel defines the methods for model operations
type IModel[T any] interface {
	FindAll(where interface{}) ([]T, error)
	FindOne(where interface{}) (*T, error)
	FindOrCreate(where interface{}, values interface{}) (*T, error)
}

// GormModel is a concrete implementation of IModel using GORM
type GormModel[T any] struct {
	DB *gorm.DB
}

// NewGormModel initializes a GormModel with a gorm.DB instance
func NewGormModel[T any](db *gorm.DB) *GormModel[T] {
	return &GormModel[T]{DB: db}
}

// FindAll retrieves all records of type T that match the given conditions
func (g *GormModel[T]) FindAll(where interface{}) ([]T, error) {
	var results []T
	err := g.DB.Where(where).Find(&results).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find all records fail")
	}
	return results, nil
}

// FindOne retrieves a single record of type T that matches the given conditions
func (g *GormModel[T]) FindOne(where interface{}) (*T, error) {
	var result T
	err := g.DB.Where(where).First(&result).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find one record fail")
	}
	return &result, nil
}

// FindOrCreate retrieves or creates a record of type T
func (g *GormModel[T]) FindOrCreate(where interface{}, values interface{}) (*T, error) {
	var result T
	err := g.DB.Where(where).Attrs(values).FirstOrCreate(&result).Error
	if err != nil {
		return nil, errors.Wrapf(err, "find or create record fail")
	}
	return &result, nil
}

// User struct represents a user model
type Student struct {
	Username string
	Email    string
}

// Example: Usage
func StudentUsage(db *gorm.DB) {
	studentModel := NewGormModel[Student](db)
	// Create user
	_, err := studentModel.FindOrCreate(map[string]interface{}{"username": "john_doe"}, map[string]interface{}{"email": "john@example.com"})
	if err != nil {
		// Handle error
		return
	}

	// Find user
	foundStudent, err := studentModel.FindOne(map[string]interface{}{"username": "john_doe"})
	if err != nil {
		// Handle error
		return
	}

	// Output found user information
	fmt.Printf("Found User: %s, Email: %s\n", foundStudent.Username, foundStudent.Email)
}

// MongoOrmModel is a concrete implementation of IModel using MongoDB
type MongoOrmModel[T any] struct {
	Collection *mongo.Collection
}

// NewMongoOrmModel initializes a MongoOrmModel with a MongoDB collection
func NewMongoOrmModel[T any](collection *mongo.Collection) *MongoOrmModel[T] {
	return &MongoOrmModel[T]{Collection: collection}
}

// FindAll retrieves all records of type T that match the given conditions
func (m *MongoOrmModel[T]) FindAll(where interface{}) ([]T, error) {
	var results []T
	cursor, err := m.Collection.Find(context.TODO(), where)
	if err != nil {
		return nil, errors.Wrap(err, "find all records fail")
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result T
		if err := cursor.Decode(&result); err != nil {
			return nil, errors.Wrap(err, "decode record fail")
		}
		results = append(results, result)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor error")
	}

	return results, nil
}

// FindOne retrieves a single record of type T that matches the given conditions
func (m *MongoOrmModel[T]) FindOne(where interface{}) (*T, error) {
	var result T
	err := m.Collection.FindOne(context.TODO(), where).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No document found
		}
		return nil, errors.Wrap(err, "find one record fail")
	}
	return &result, nil
}

// FindOrCreate retrieves or creates a record of type T
func (m *MongoOrmModel[T]) FindOrCreate(where interface{}, values interface{}) (*T, error) {
	var result T
	err := m.Collection.FindOne(context.TODO(), where).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Document not found, create it
			_, err := m.Collection.InsertOne(context.TODO(), values)
			if err != nil {
				return nil, errors.Wrap(err, "insert record fail")
			}
			// Return the newly created document
			err = m.Collection.FindOne(context.TODO(), where).Decode(&result)
			if err != nil {
				return nil, errors.Wrap(err, "find created record fail")
			}
		} else {
			return nil, errors.Wrap(err, "find record fail")
		}
	}
	return &result, nil
}

// Teacher struct represents a Teacher model for MongoDB
type Teacher struct {
	Username string `bson:"username"`
	Email    string `bson:"email"`
}

// Example: Usage
func ExampleMongoUsage(collection *mongo.Collection) {
	// Initialize MongoOrmModel for Teacher
	teacherModel := NewMongoOrmModel[Teacher](collection)

	// Create a new user
	newTeacher := Teacher{
		Username: "john_doe",
		Email:    "john@example.com",
	}
	_, err := teacherModel.FindOrCreate(map[string]interface{}{"username": newTeacher.Username}, newTeacher)
	if err != nil {
		// Handle error
		return
	}

	// Find user
	foundTeacher, err := teacherModel.FindOne(map[string]interface{}{"username": "john_doe"})
	if err != nil {
		// Handle error
		return
	}

	// Output found user information
	if foundTeacher != nil {
		fmt.Printf("Found foundTeacher: %s, Email: %s\n", foundTeacher.Username, foundTeacher.Email)
	} else {
		fmt.Println("foundTeacher not found")
	}
}
