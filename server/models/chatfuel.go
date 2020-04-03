package models

import e "gin-vegan-map-api/server/enum"

// Text -
type Text struct {
	Text string `json:"text"`
}

// Messages -
type Messages struct {
	Messages interface{} `json:"messages"`
}

// Button -
type Button struct {
	Type  string `json:"type"`
	URL   string `json:"url"`
	Title string `json:"title"`
}

// GalleryElements -
type GalleryElements struct {
	Title    string   `json:"title"`
	ImageURL string   `json:"image_url"`
	Subtitle string   `json:"subtitle"`
	Buttons  []Button `json:"buttons"`
}

// Payload -
type Payload struct {
	TemplateType     string            `json:"template_type"`
	ImageAspectRatio string            `json:"ImageAspectRatio"`
	Elements         []GalleryElements `json:"elements"`
}

// Attachment -
type Attachment struct {
	Type    string  `json:"type"`
	Payload Payload `json:"payload"`
}

// Attachments -
type Attachments struct {
	Attachments Attachment `json:"attachment"`
}

// GetRestaurantByFriedRegional -
func GetRestaurantByFriedRegional(regional string) (Messages, error) {
	restaurants, err := GetRestaurantsByFriedRegional(regional)
	if err != nil {
		return Messages{}, err
	}

	var elements []GalleryElements
	button := Button{
		Type:  e.ButtonTypeWeb,
		URL:   e.FbPageURL,
		Title: e.ClickToWebPage,
	}

	// chatfuel up to 10 rows, 0-9
	for i, restaurant := range restaurants {
		if i > 9 {
			break
		}

		element := GalleryElements{
			Title:    restaurant.Name,
			ImageURL: e.DefaultImageURL,
			Subtitle: restaurant.Address,
			Buttons:  []Button{button},
		}
		elements = append(elements, element)
	}

	// Create payload
	payload := Payload{
		TemplateType:     e.PayloadTemplateTypeGeneric,
		ImageAspectRatio: e.ImageAspectRatio,
		Elements:         elements,
	}

	// Create attachment
	attachment := Attachments{
		Attachments: Attachment{
			Type:    e.AttachmentTypeTemplate,
			Payload: payload,
		},
	}

	return Messages{
		[]Attachments{attachment},
	}, nil
}

func getGalleryMessages(pics []Pics) (Messages, error) {
	var elements []GalleryElements

	// chatfuel up to 10 rows, 0-9
	for i, pic := range pics {
		if i > 9 {
			break
		}

		button := Button{
			Type:  e.ButtonTypeWeb,
			URL:   pic.URL,
			Title: e.ClickToWebPage,
		}

		element := GalleryElements{
			Title:    pic.Name,
			ImageURL: pic.URL,
			Subtitle: pic.Regional,
			Buttons:  []Button{button},
		}
		elements = append(elements, element)
	}

	// Create payload
	payload := Payload{
		TemplateType:     e.PayloadTemplateTypeGeneric,
		ImageAspectRatio: e.ImageAspectRatio,
		Elements:         elements,
	}

	// Create attachment
	attachment := Attachments{
		Attachments: Attachment{
			Type:    e.AttachmentTypeTemplate,
			Payload: payload,
		},
	}

	return Messages{
		[]Attachments{attachment},
	}, nil
}

// GetGalleryMessagesByTypeAndRegional -
func GetGalleryMessagesByTypeAndRegional(typeName string, regional string) (Messages, error) {
	pics, err := GetPicsByTypeAndRegional(typeName, regional)
	if err != nil {
		return Messages{}, err
	}

	return getGalleryMessages(pics)
}