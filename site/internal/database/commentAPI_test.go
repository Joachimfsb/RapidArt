package database

/*
func TestAddCommentToPost(t *testing.T) {

	// Declare expectations
	//mock.ExpectCommit()
	user := test.GenTestUser()

	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)

	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)
	comment := test.GenComment(user.UserId, post.PostId)

	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(comment.UserId, comment.PostId, comment.Message, comment.CreationDateTime).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := AddCommentToPost(comment); err != nil {
		t.Fatal("Got error trying to add report: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

// inspirert av TestGetAllReportsForPost(...)
func TestGetAllCommentsFromPost(t *testing.T) {
	// Generate test data
	user := test.GenTestUser()
	gallery := test.GenBasisGallery()
	canvas := test.GenBasisCanvas(gallery.BasisGalleryId)
	post, _ := test.GenTestPost(user.UserId, canvas.BasisCanvasId, false)
	comment := test.GenReport(user.UserId, post.PostId)

	// Expect the INSERT query in NewReport
	mock.ExpectExec(`^INSERT INTO Report`).WithArgs(comment.UserId, comment.PostId, comment.Message, comment.CreationDateTime).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Call NewReport to execute the INSERT expectation
	if err := NewReport(comment); err != nil {
		t.Fatal("Error inserting report: " + err.Error())
	}

	// Now expect the SELECT query for GetAllReportsForPost
	rows := sqlmock.NewRows([]string{"UserId", "PostId", "Message", "CreationDateTime"}).
		AddRow(comment.UserId, comment.PostId, comment.Message, comment.CreationDateTime)

	mock.ExpectQuery(`^SELECT \* FROM Report WHERE PostId = ?`).WithArgs(comment.PostId).WillReturnRows(rows)

	// Call GetAllReportsForPost to execute the SELECT expectation
	reportAmount, err := GetAllReportsForPost(comment.PostId)
	if err != nil {
		t.Fatal("Got error trying to count reports: " + err.Error())
	}

	// Check that the returned report amount is as expected
	if len(reportAmount) != 1 {
		t.Fatal("Wrong amount of reports: expected 1, got", len(reportAmount))
	}

	// Ensure all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
*/
