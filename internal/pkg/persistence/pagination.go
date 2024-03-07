package persistence

import (
	"fmt"
	"neversitup-test-template/internal/pkg/models"
)

type PaginationRepository struct{}

var paginationRepository *PaginationRepository

func Pagination() *PaginationRepository {
	if paginationRepository == nil {
		paginationRepository = &PaginationRepository{}
	}
	return paginationRepository
}

func (r *PaginationRepository) SetPaging(Data interface{}, Header interface{}, CurrentURL string, CurrentPage int, PerPage int, RowCount int) models.Pagination {
	var result models.Pagination
	result.Data = Data
	result.Header = Header
	Meta := models.Meta{
		CurrentPage: CurrentPage,
		Path:        CurrentURL,
		PerPage:     PerPage,
		Total:       RowCount,
	}
	Meta.LastPage = (Meta.Total-1)/Meta.PerPage + 1
	if Meta.CurrentPage < 6 {
		Meta.From = 1
		if Meta.LastPage >= 10 {
			Meta.To = 10
		} else {
			Meta.To = Meta.LastPage
		}
	} else if Meta.CurrentPage > Meta.LastPage-6 {
		if Meta.LastPage > 10 {
			Meta.From = Meta.LastPage - 9
		} else {
			Meta.From = 1
		}
		Meta.To = Meta.LastPage
	} else {
		Meta.From = Meta.CurrentPage - 4
		Meta.To = Meta.CurrentPage + 5
	}
	//Meta.Path
	//Meta.PerPage
	//Meta.Total
	if Meta.CurrentPage == 1 {
		link := models.LinkPage{
			Label:  "&laquo; Previous",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)
	} else {
		link := models.LinkPage{
			Url:    r.CondSprintf(Meta.Path, Meta.CurrentPage-1, PerPage),
			Label:  "&laquo; Previous",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)
	}
	if Meta.From != 1 {
		linkFirst := models.LinkPage{
			Url:    r.CondSprintf(Meta.Path, 1, PerPage),
			Label:  "First",
			Active: false,
		}
		Meta.Links = append(Meta.Links, linkFirst)

		link := models.LinkPage{
			Label:  "...",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)
	}
	for i := Meta.From; i <= Meta.To; i++ {
		link := models.LinkPage{
			Url:    r.CondSprintf(Meta.Path, i, PerPage),
			Label:  fmt.Sprintf("%d", i),
			Active: i == Meta.CurrentPage,
		}
		Meta.Links = append(Meta.Links, link)
	}
	if Meta.To != Meta.LastPage {
		link := models.LinkPage{
			Label:  "...",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)

		linkLast := models.LinkPage{
			Url:    r.CondSprintf(Meta.Path, Meta.LastPage, PerPage),
			Label:  "Last",
			Active: false,
		}
		Meta.Links = append(Meta.Links, linkLast)
	}
	if Meta.CurrentPage == Meta.LastPage {
		link := models.LinkPage{
			Label:  "Next &raquo;",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)
	} else {
		link := models.LinkPage{
			Url:    r.CondSprintf(Meta.Path, Meta.CurrentPage+1, PerPage),
			Label:  "Next &raquo;",
			Active: false,
		}
		Meta.Links = append(Meta.Links, link)
	}
	result.Meta = Meta
	result.Links.First = r.CondSprintf(Meta.Path, 1, PerPage)
	result.Links.Last = r.CondSprintf(Meta.Path, Meta.LastPage, PerPage)
	if Meta.CurrentPage != 1 {
		result.Links.Prev = r.CondSprintf(Meta.Path, Meta.CurrentPage-1, PerPage)
	}
	if Meta.CurrentPage != Meta.LastPage {
		result.Links.Next = r.CondSprintf(Meta.Path, Meta.CurrentPage+1, PerPage)
	}
	result.Meta.Path = r.CondSprintf(CurrentURL, CurrentPage, PerPage)

	return result
}

func (r *PaginationRepository) CondSprintf(format string, v ...interface{}) string {
	v = append(v, "")
	format += fmt.Sprint("%[", len(v), "]s")
	return fmt.Sprintf(format, v...)
}
