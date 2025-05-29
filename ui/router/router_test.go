package router

import "testing"

func TestRouter_trimHead(t *testing.T) {
	tests := []struct {
		name    string
		maxSize int
		amount  int
	}{
		{
			maxSize: 2,
			amount:  1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newRouter(tt.maxSize)
			router.navigate("example_page_1")
			router.navigate("example_page_2")
			sizeBefore := router.size
			router.trimHead(tt.amount)

			if sizeBefore != 2 || router.size != 1 ||
				router.location.path != "example_page_2" {
				t.Errorf(
					"router.trimHead() = size %v, want size %v, size before %v, want size before %v, current page name %v",
					router.size,
					1,
					sizeBefore,
					2,
					router.location.path,
				)
			}
		})
	}
}

func TestRouter_Navigate(t *testing.T) {
	tests := []struct {
		name               string
		maxSize            int
		pages              []string
		wantEndCurrentPage string
	}{
		{
			maxSize:            4,
			pages:              []string{"example_page_1", "example_page_2", "example_page_3"},
			wantEndCurrentPage: "example_page_3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newRouter(tt.maxSize)

			for _, page := range tt.pages {
				router.navigate(page)
			}

			if router.location.path != tt.wantEndCurrentPage {
				t.Errorf(
					"router.Navigate() = currentPage = %v, want %v",
					router.location.path,
					tt.wantEndCurrentPage,
				)
			}
		})
	}
}

func TestRouter_Forward(t *testing.T) {
	tests := []struct {
		name     string
		maxSize  int
		pages    []string
		back     int
		forward  int
		wantPage string
	}{
		{
			maxSize:  4,
			pages:    []string{"example_page_1", "example_page_2", "example_page_3"},
			back:     2,
			forward:  1,
			wantPage: "example_page_2",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newRouter(tt.maxSize)

			for _, page := range tt.pages {
				router.navigate(page)
			}

			for range tt.back {
				router.back()
			}

			for range tt.forward {
				router.forward()
			}

			if router.location.path != tt.wantPage {
				t.Errorf(
					"router.Forward() = currentPage = %v, want %v",
					router.location.path,
					tt.wantPage,
				)
			}
		})
	}
}

func TestRouter_Back(t *testing.T) {
	tests := []struct {
		name     string
		maxSize  int
		pages    []string
		back     int
		wantPage string
	}{
		{
			maxSize:  4,
			pages:    []string{"example_page_1", "example_page_2", "example_page_3"},
			back:     2,
			wantPage: "example_page_1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newRouter(tt.maxSize)
			for _, page := range tt.pages {
				router.navigate(page)
			}

			for range tt.back {
				router.back()
			}

			if router.location.path != tt.wantPage {
				t.Errorf(
					"router.Forward() = currentPage = %v, want %v",
					router.location.path,
					tt.wantPage,
				)
			}
		})
	}
}

func TestRouter_Clear(t *testing.T) {
	tests := []struct {
		name    string
		maxSize int
		pages   []string
	}{
		{
			maxSize: 4,
			pages:   []string{"example_page_1", "example_page_2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := newRouter(tt.maxSize)

			for _, page := range tt.pages {
				router.navigate(page)
			}

			router.clear()

			if router.size != 0 {
				t.Error("Router size > 0")
			}

			if router.location != nil {
				t.Error("router current page not nil")
			}

			if router.head != nil {
				t.Error("router head not nil")
			}
		})
	}
}
