package isan

type ISAN struct {
	recordType   string
	title        string
	workType     string
	releaseDate  string
	director     string
	durationMin  string
	originalLang string
	isan         string
}

func CreateNewISAN(isan_num string, recordtype string,
	title string, worktype string, releasedate string,
	director string, durationMin string, originalLang string) ISAN {

	var newisan ISAN

	newisan.isan = isan_num             // TODO SETTERS AND CHECKERS
	newisan.recordType = recordtype     // TODO SETTERS AND CHECKERS
	newisan.title = title               // TODO SETTERS AND CHECKERS
	newisan.workType = worktype         // TODO SETTERS AND CHECKERS
	newisan.releaseDate = releasedate   // TODO SETTERS AND CHECKERS
	newisan.director = director         // TODO SETTERS AND CHECKERS
	newisan.durationMin = durationMin   // TODO SETTERS AND CHECKERS
	newisan.originalLang = originalLang // TODO SETTERS AND CHECKERS

	return newisan
}

func (i *ISAN) GetIsanVals() (isan_num, record_type, title, work_type, release_date, director, duration_min, original_language string) {
	isan_num = i.isan
	record_type = i.recordType
	title = i.title
	work_type = i.workType
	release_date = i.releaseDate
	director = i.director
	duration_min = i.durationMin
	original_language = i.originalLang

	return
}
