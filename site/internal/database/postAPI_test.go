package database

import (
	"rapidart/internal/models"
	"rapidart/test"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestShouldNotGetPostsByUserId(t *testing.T) {
	// Test data
	uid := 1
	var limit uint = 5

	// Declare expectations
	mock.ExpectQuery("SELECT").
		WithArgs(uid, limit).
		WillReturnRows(sqlmock.NewRows(nil))

	// Function call
	ret, err := GetPostsByUserId(uid, "creationDateTimeAsc", limit)
	if err != nil {
		t.Fatal("Error returned from function call!")
	}
	if ret != nil {
		t.Fatal("Got data when we shouldn't have!")
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestAddPost(t *testing.T) {
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)

	// Declare db expectations
	mock.ExpectExec(`^INSERT (.+)`).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if _, err := AddPost(post); err != nil {
		t.Fatal("Got error trying to create user: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestGetPostById(t *testing.T) {

	// Test data
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)

	// Declare expectations
	mock.ExpectQuery("^SELECT").WithArgs(post.PostId).WillReturnRows(GenRows([]models.Post{post}))

	// Function call
	_, err := GetPostById(post.PostId)
	if err != nil {
		t.Fatal("Got error trying to get post by id: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

// Laget ved hjelp av KI/AI
// https://chatgpt.com/share/6720aac0-27e0-800c-92e8-1ce43ddcbb0e
func TestGetPostsWithLikeCountSortedByMostLikes(t *testing.T) {
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	var elevenPosts []models.PostExtended

	var i = 0
	//creates 11 posts
	for i <= 10 {
		_, post := test.GenTestPost(user.UserId, canvas.BasisCanvasId, true)
		elevenPosts = append(elevenPosts, post)
		i += 1
	}

	// Prepare the mock result set
	rows := sqlmock.NewRows([]string{
		"PostId", "UserId", "BasisCanvasId", "Image", "Caption", "TimeSpentDrawing", "CreationDateTime", "LikeCount",
	})
	for _, post := range elevenPosts {
		rows.AddRow(post.PostId, post.UserId, post.BasisCanvasId, post.Image, post.Caption, post.TimeSpentDrawing, post.CreationDateTime, post.LikeCount)
	}

	// Declare expectations
	mock.ExpectQuery("^SELECT p.PostId, p.UserId, p.BasisCanvasId, p.Image, p.Caption, p.TimeSpentDrawing, p.CreationDateTime, COUNT").
		WithArgs(10).
		WillReturnRows(rows)

	// Call the function to test
	result, err := GetPostsWithLikeCountSortedByMostLikes(10)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	// Verify the result length and LikeCount order
	if len(result) != 10 {
		t.Fatalf("expected 10 posts, but got: %d", len(result))
	}
	for i = 0; i < len(result)-1; i++ {
		if result[i].LikeCount < result[i+1].LikeCount {
			t.Fatalf("expected posts to be sorted by LikeCount in descending order")
		}
	}

	// Ensure all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetRecentPostsWithLikeCount(t *testing.T) {
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	var tenPosts []models.PostExtended

	var i = 0
	//creates 10 posts
	for i < 10 {
		_, post := test.GenTestPost(user.UserId, canvas.BasisCanvasId, true)
		tenPosts = append(tenPosts, post)
		i += 1
	}

	// Prepare the mock result set
	rows := sqlmock.NewRows([]string{
		"PostId", "UserId", "BasisCanvasId", "Image", "Caption", "TimeSpentDrawing", "CreationDateTime", "Active", "LikeCount",
	})
	for _, post := range tenPosts {
		rows.AddRow(post.PostId, post.UserId, post.BasisCanvasId, post.Image, post.Caption, post.TimeSpentDrawing, post.CreationDateTime, post.Active, post.LikeCount)
	}

	// Declare expectations
	mock.ExpectQuery("^SELECT").
		WithArgs(true, 10).
		WillReturnRows(rows)

	// Call the function to test
	result, err := GetRecentPostsWithLikes(10, true)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	// Verify the result length and CreationDateTime order
	if len(result) != 10 {
		t.Fatalf("expected 10 posts, but got: %d", len(result))
	}

	// Ensure all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}

func TestGetRecentFollowsPostsWithLikeCount(t *testing.T) {
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	var tenPosts []models.PostExtended

	var i = 0
	//creates 11 posts
	for i < 10 {
		_, post := test.GenTestPost(user.UserId, canvas.BasisCanvasId, true)
		tenPosts = append(tenPosts, post)
		i += 1
	}

	// Prepare the mock result set
	rows := sqlmock.NewRows([]string{
		"PostId", "UserId", "BasisCanvasId", "Image", "Caption", "TimeSpentDrawing", "CreationDateTime", "Active", "LikeCount",
	})
	for _, post := range tenPosts {
		rows.AddRow(post.PostId, post.UserId, post.BasisCanvasId, post.Image, post.Caption, post.TimeSpentDrawing, post.CreationDateTime, post.Active, post.LikeCount)
	}

	// Declare expectations
	mock.ExpectQuery("^SELECT").
		WithArgs(true, 10).
		WillReturnRows(rows)

	// Call the function to test
	result, err := GetUsersFollowsRecentPostsWithLikes(user.UserId, 10, true)
	if err != nil {
		t.Fatalf("expected no error, but got: %v", err)
	}

	// Verify the result length and CreationDateTime order
	if len(result) != 10 {
		t.Fatalf("expected 10 posts, but got: %d", len(result))
	}

	// Ensure all expectations were met
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("there were unfulfilled expectations: %v", err)
	}
}
