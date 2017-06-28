package findface

import (
	"context"
)

type FaceVerifyOptions struct {
	// The first image external URL
	FirstPhoto string `json:"photo1"`

	// Bounding boxes for the faces on the first photo.
	FirstBoundingBox *BoundingBox `json:"bbox1"`

	// The second image external URL
	SecondPhoto string `json:"photo2"`

	// Bounding boxes for the faces on the second photo.
	SecondBoundingBox *BoundingBox `json:"bbox2"`

	// Specifies behavior in case if multiple faces are detected on the photo; one of:
	// `reject` return an error and a list of faces if more than one face is detected on the provided photo
	// `biggest` (default) search using the biggest face on the photo
	// `all` search for each face found on the photo.
	MulitFaceSelector string `json:"mf_selector"`

	// [optional]: one of "strict", "medium", "low" [default], "none" or a value between 0 and 1
	// Example: "0.75"
	Threshold string `json:"threshold"`
}

type FaceVerifyResultResponse struct {
	FindFaceResponse
	Results []*FaceVerifyResult `json:"results"`
}

type FaceVerifyResult struct {
	// Bounding boxes for the faces on the first photo.
	FirstBoundingBox *BoundingBox `json:"bbox1"`
	// Bounding boxes for the faces on the second photo.
	SecondBoundingBox *BoundingBox `json:"bbox2"`

	Confidence float32 `json:"confidence"`
	Verified   bool    `json:"verified"`
}

func (s *FacesService) Verify(ctx context.Context, opt *FaceVerifyOptions) (*FaceVerifyResultResponse, error) {
	req, err := s.client.NewRequest("POST", "/verify", opt)
	if err != nil {
		return nil, err
	}

	var result *FaceVerifyResultResponse
	resp, err := s.client.Do(ctx, req, &result)
	result.Response = resp
	return result, err
}