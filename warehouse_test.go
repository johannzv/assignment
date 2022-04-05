package main

import (
	"reflect"
	"testing"
)

func TestWarehouse_GetProducts(t *testing.T) {
	type fields struct {
		inventory []Article
		product   []warehouseProduct
	}
	tests := []struct {
		name   string
		fields fields
		want   []warehouseProduct
	}{
		{
			name: "Empty warehouse, expect empty result",
			fields: fields{
				inventory: []Article{},
				product:   []warehouseProduct{},
			},
			want: []warehouseProduct{},
		},
		{
			name: "Warehouse contains table and article needed for a table",
			fields: fields{
				inventory: []Article{
					{ArtId: "1", Name: "leg", Stock: 4},
					{ArtId: "2", Name: "screw", Stock: 8},
					{ArtId: "4", Name: "table top", Stock: 1},
				},
				product: []warehouseProduct{
					{
						ProductId: "0",
						Name:      "Dining Table",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "4",
								AmountOf: 1,
							},
						},
					},
				},
			},
			want: []warehouseProduct{{
				Name:      "Dining Table",
				ProductId: "0",
				Articles: []ContainArticle{
					{
						ArtId:    "1",
						AmountOf: 4,
					},
					{
						ArtId:    "2",
						AmountOf: 8,
					},
					{
						ArtId:    "4",
						AmountOf: 1,
					},
				},
				Quantity: 1,
			}},
		},
		{
			name: "Warehouse contains table and article needed for a table and a chair, but not for both",
			fields: fields{
				inventory: []Article{
					{ArtId: "1", Name: "leg", Stock: 4},
					{ArtId: "2", Name: "screw", Stock: 8},
					{ArtId: "3", Name: "seat", Stock: 1},
					{ArtId: "4", Name: "table top", Stock: 1},
				},
				product: []warehouseProduct{
					{
						ProductId: "0",
						Name:      "Dining Table",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "4",
								AmountOf: 1,
							},
						},
					},
					{
						ProductId: "1",
						Name:      "Dining Chair",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "3",
								AmountOf: 1,
							},
						},
					},
				},
			},
			want: []warehouseProduct{{
				ProductId: "0",
				Name:      "Dining Table",
				Articles: []ContainArticle{
					{
						ArtId:    "1",
						AmountOf: 4,
					},
					{
						ArtId:    "2",
						AmountOf: 8,
					},
					{
						ArtId:    "4",
						AmountOf: 1,
					},
				},
				Quantity: 1,
			}, {
				ProductId: "1",
				Name:      "Dining Chair",
				Articles: []ContainArticle{
					{
						ArtId:    "1",
						AmountOf: 4,
					},
					{
						ArtId:    "2",
						AmountOf: 8,
					},
					{
						ArtId:    "3",
						AmountOf: 1,
					},
				},
				Quantity: 1,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &Warehouse{
				inventory: tt.fields.inventory,
				product:   tt.fields.product,
			}
			if got := wh.GetProducts(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProducts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarehouse_SellProduct(t *testing.T) {
	type fields struct {
		inventory []Article
		product   []warehouseProduct
	}
	type args struct {
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Sell product which is in stock",
			fields: fields{
				inventory: []Article{
					{ArtId: "1", Name: "leg", Stock: 4},
					{ArtId: "2", Name: "screw", Stock: 8},
					{ArtId: "4", Name: "table top", Stock: 1},
				},
				product: []warehouseProduct{
					{
						ProductId: "0",
						Name:      "Dining Table",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "4",
								AmountOf: 1,
							},
						},
					},
				},
			},
			args: args{
				productId: "0",
			},
			wantErr: false,
		},
		{
			name: "Sell product (ProductId 9) which does not exist",
			fields: fields{
				inventory: []Article{
					{ArtId: "1", Name: "leg", Stock: 4},
					{ArtId: "2", Name: "screw", Stock: 8},
					{ArtId: "4", Name: "table top", Stock: 1},
				},
				product: []warehouseProduct{
					{
						ProductId: "0",
						Name:      "Dining Table",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "4",
								AmountOf: 1,
							},
						},
					},
				},
			},
			args: args{
				productId: "9",
			},
			wantErr: true,
		},
		{
			name: "Sell product which exist, but leg (artId 1) is out of stock",
			fields: fields{
				inventory: []Article{
					{ArtId: "1", Name: "leg", Stock: 0},
					{ArtId: "2", Name: "screw", Stock: 8},
					{ArtId: "4", Name: "table top", Stock: 1},
				},
				product: []warehouseProduct{
					{
						ProductId: "0",
						Name:      "Dining Table",
						Articles: []ContainArticle{
							{
								ArtId:    "1",
								AmountOf: 4,
							},
							{
								ArtId:    "2",
								AmountOf: 8,
							},
							{
								ArtId:    "4",
								AmountOf: 1,
							},
						},
					},
				},
			},
			args: args{
				productId: "0",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &Warehouse{
				inventory: tt.fields.inventory,
				product:   tt.fields.product,
			}
			if err := wh.SellProduct(tt.args.productId); (err != nil) != tt.wantErr {
				t.Errorf("SellProduct() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestWarehouse_GetStockForArticleByArticleId(t *testing.T) {
	type fields struct {
		inventory []Article
		product   []warehouseProduct
	}
	type args struct {
		artId string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "Article is not in stock",
			fields: fields{
				inventory: []Article{},
				product:   []warehouseProduct{},
			},
			args: args{
				artId: "1",
			},
			want: 0,
		},
		{
			name: "Article is in stock",
			fields: fields{
				inventory: []Article{{
					ArtId: "1",
					Name:  "legs",
					Stock: 3,
				}},
				product: []warehouseProduct{},
			},
			args: args{
				artId: "1",
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &Warehouse{
				inventory: tt.fields.inventory,
				product:   tt.fields.product,
			}
			if got := wh.GetStockForArticleByArticleId(tt.args.artId); got != tt.want {
				t.Errorf("GetStockForArticleByArticleId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWarehouse_GetProductByProductId(t *testing.T) {
	type fields struct {
		inventory []Article
		product   []warehouseProduct
	}
	type args struct {
		productId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    warehouseProduct
		wantErr bool
	}{
		{
			name: "product is not in warehouse",
			fields: fields{
				inventory: []Article{},
				product:   []warehouseProduct{},
			},
			args:    args{productId: "2"},
			want:    warehouseProduct{},
			wantErr: true,
		},
		{
			name: "product is in warehouse",
			fields: fields{
				product: []warehouseProduct{{
					Name:      "Chair",
					Articles:  nil,
					Quantity:  0,
					ProductId: "1",
				}},
			},
			want: warehouseProduct{
				Name:      "Chair",
				Articles:  nil,
				Quantity:  0,
				ProductId: "1",
			},
			args:    args{productId: "1"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wh := &Warehouse{
				inventory: tt.fields.inventory,
				product:   tt.fields.product,
			}
			got, err := wh.GetProductByProductId(tt.args.productId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetProductByProductId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProductByProductId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getMin(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "find min",
			args: args{array: []int{1, 2, 3}},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMin(tt.args.array); got != tt.want {
				t.Errorf("getMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
