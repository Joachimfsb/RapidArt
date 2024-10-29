package database

/*
func TestGetBasisGalleryById(t *testing.T) {

	// Test data
	gallery := test.GenBasisGallery()

	// Declare expectations
	mock.ExpectQuery("^SELECT").WithArgs(gallery.BasisGalleryId).WillReturnRows(GenRows([]models.BasisGallery{gallery}))

	// Function call
	_, err := GetBasisGalleryById(gallery.BasisGalleryId)
	if err != nil {
		t.Fatal("Got error trying to get gallery by id: " + err.Error())
	}
	// Check if expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}

func TestAddGallery(t *testing.T) {

	// Declare expectations
	//mock.ExpectCommit()
	gallery := test.GenBasisGallery()

	//mock.ExpectBegin()
	mock.ExpectExec(`^INSERT (.+)`).WithArgs(gallery.StartDateTime, gallery.EndDateTime).WillReturnResult(sqlmock.NewResult(1, 1))

	// Function call
	if err := AddGallery(gallery); err != nil {
		t.Fatal("Got error trying to add report: " + err.Error())
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal("Some expectations were not met: " + err.Error())
	}
}
*/
