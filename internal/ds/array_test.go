package ds

import (
	"reflect"
	"sort"
	"testing"
)

func TestJSONArray_String(t *testing.T) {
	type fields struct {
		Key      string
		Children []JSONElement
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Array Test Only Bools",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONPrimitive{
						Ptype: Bool,
						Key:   "testbool",
					},
					&JSONPrimitive{
						Ptype: Bool,
						Key:   "testbool2",
					},
				},
			},
			want: "Testarray []bool `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Only Strings",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONPrimitive{
						Ptype: String,
						Key:   "teststring1",
					},
					&JSONPrimitive{
						Ptype: String,
						Key:   "teststring2",
					},
				},
			},
			want: "Testarray []string `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Only Ints",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONPrimitive{
						Ptype: Int,
						Key:   "testint1",
					},
					&JSONPrimitive{
						Ptype: Int,
						Key:   "testint2",
					},
				},
			},
			want: "Testarray []int `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Arrays",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONArray{
						Key: "Doesn't matter",
						Children: []JSONElement{
							&JSONPrimitive{
								Ptype: Bool,
								Key:   "testbool",
							},
							&JSONPrimitive{
								Ptype: Int,
								Key:   "testint",
							},
						}},
					&JSONArray{
						Key: "Doesn't matter aswell",
						Children: []JSONElement{
							&JSONPrimitive{
								Ptype: Bool,
								Key:   "testbool",
							},
							&JSONPrimitive{
								Ptype: Int,
								Key:   "testint",
							},
						},
					},
				},
			},
			want: "Testarray [][]interface{} `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Different Objects",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONObject{
						Key: "Intstringobj",
						Children: []JSONElement{
							&JSONPrimitive{
								Ptype: Int,
								Key:   "testint",
							},
							&JSONPrimitive{
								Ptype: String,
								Key:   "teststring",
							},
						},
					},
					&JSONObject{
						Key: "Boolobj",
						Children: []JSONElement{
							&JSONPrimitive{
								Ptype: Bool,
								Key:   "testbool",
							},
						},
					},
				},
			},
			want: "Testarray []struct{\n" +
				"Testint int `json:\"testint\"`\n" +
				"Teststring string `json:\"teststring\"`\n" +
				"Testbool bool `json:\"testbool\"`\n" +
				"} `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Different Types Primitive",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONPrimitive{
						Ptype: Bool,
						Key:   "testbool",
					},
					&JSONPrimitive{
						Ptype: Int,
						Key:   "testint",
					},
				},
			},
			want: "Testarray []interface{} `json:\"testarray\"`\n",
		},
		{
			name: "Array Test Different Types Mixed",
			fields: fields{
				Key: "testarray",
				Children: []JSONElement{
					&JSONPrimitive{
						Ptype: Bool,
						Key:   "testbool",
					},
					&JSONObject{
						Key: "Doesn't matter",
						Children: []JSONElement{
							&JSONPrimitive{
								Ptype: Int,
								Key:   "testint",
							},
							&JSONPrimitive{
								Ptype: String,
								Key:   "teststring",
							},
						},
					},
				},
			},
			want: "Testarray []interface{} `json:\"testarray\"`\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jp := &JSONArray{
				Key:      tt.fields.Key,
				Children: tt.fields.Children,
			}
			if got := jp.String(); got != tt.want {
				t.Errorf("JSONArray.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_listChildrenTypes(t *testing.T) {
	type args struct {
		c []JSONElement
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Different Primitives",
			args: args{
				[]JSONElement{
					&JSONPrimitive{
						Ptype: Bool,
						Key:   "Testbool",
					},
					&JSONPrimitive{
						Ptype: String,
						Key:   "Teststring",
					},
					&JSONPrimitive{
						Ptype: Float,
						Key:   "Testfloat",
					},
					&JSONPrimitive{
						Ptype: Int,
						Key:   "Testint",
					},
				},
			},
			want: []string{"bool", "int", "float64", "string"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := listChildrenTypes(tt.args.c)
			want := tt.want
			sort.Slice(want, func(i, j int) bool {
				return want[i] < want[j]
			})
			sort.Slice(got, func(i, j int) bool {
				return got[i] < got[j]
			})
			if !reflect.DeepEqual(got, want) {
				t.Errorf("listChildrenTypes() = %v, want %v", got, tt.want)
			}
		})
	}
}
