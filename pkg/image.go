package MD_Builder

func (image *Image) New_Image() *Image {
	return &Image{
		Source: "",
	}
}

func (image *Image) Set_Source(source string) {
	image.Source = source
}
