package builder

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	. "sdui.builder/models"
	. "sdui.builder/pathfinders"
)

func BuildLayoutDTO(data Data, layout Layout) LayoutDTO {
	data.ParseRaw()
	var dto LayoutDTO
	setBody(&dto, &layout)
	parsePages(&dto, &layout, &data)
	clearDTORepetitions(&dto)
	return dto
}

func setBody(dto *LayoutDTO, layout *Layout) {
	dto.Code = layout.Code
}

func parsePages(dto *LayoutDTO, layout *Layout, data *Data) {
	newPages := make([]PageDTO, 0)
	var pages []Page = layout.Pages

	sort.Slice(pages, func(i, j int) bool {
		return pages[i].Order < pages[j].Order
	})

	for _, page := range layout.Pages {
		newPage := PageDTO{
			ID:          page.ID,
			SectionsIDs: parseSections(dto, page.Sections, data),
		}
		if len(*newPage.SectionsIDs) == 0 {
			continue
		}
		newPages = append(newPages, newPage)
	}

	dto.Pages = append(dto.Pages, newPages...)
}

func parseSections(dto *LayoutDTO, sections []Section, data *Data) *[]string {
	sectionsID := make([]string, 0)
	newSections := make([]SectionDTO, 0)

	sort.Slice(sections, func(i, j int) bool {
		return sections[i].Order < sections[j].Order
	})

	for _, section := range sections {
		newSection := SectionDTO{
			ID:          section.ID,
			ContentsIDs: parseContents(dto, section.Contents, data),
		}
		if len(*newSection.ContentsIDs) == 0 {
			continue
		}
		newSections = append(newSections, newSection)
		sectionsID = append(sectionsID, newSection.ID)
	}

	dto.Sections = append(dto.Sections, newSections...)
	return &sectionsID
}

func parseContents(dto *LayoutDTO, contents []Content, data *Data) *[]string {
	contentsID := make([]string, 0)
	newContents := make([]ContentDTO, 0)

	sort.Slice(contents, func(i, j int) bool {
		return contents[i].Order < contents[j].Order
	})

	for _, content := range contents {
		newContent := ContentDTO{
			ID:    content.ID,
			Value: buildValue(content.Value, data.GetParsed()),
		}
		if strings.ReplaceAll(newContent.Value, " ", "") == "" {
			continue
		}
		newContents = append(newContents, newContent)
		contentsID = append(contentsID, newContent.ID)
	}

	dto.Contents = append(dto.Contents, newContents...)
	return &contentsID
}

func buildValue(path string, data map[string]interface{}) string {
	if path, err := validatePath(path); err != nil {
		return path
	} else {
		pf := &Pathfinder{
			Path: path,
		}
		result := pf.Eval(data, path)
		return fmt.Sprint(result)
	}
}

func validatePath(path string) (string, error) {
	if strings.Contains(path, "data/") {
		return strings.Replace(path, "data/", "/", 1), nil
	} else {
		return path, errors.New("Invalid Path")
	}
}

func clearDTORepetitions(layoutDTO *LayoutDTO) {
	layoutDTO.Contents = clearSliceRepetition(layoutDTO.Contents, func(c ContentDTO) string { return c.ID })
	layoutDTO.Sections = clearSliceRepetition(layoutDTO.Sections, func(s SectionDTO) string { return s.ID })
	layoutDTO.Pages = clearSliceRepetition(layoutDTO.Pages, func(p PageDTO) string { return p.ID })
}

func clearSliceRepetition[T any, I comparable](slice []T, identifier func(T) I) []T {
	uniqueMap := make(map[I]T)
	for _, item := range slice {
		if _, ok := uniqueMap[identifier(item)]; !ok {
			uniqueMap[identifier(item)] = item
		}
	}

	result := make([]T, 0, len(uniqueMap))
	for _, item := range uniqueMap {
		result = append(result, item)
	}
	return result
}
