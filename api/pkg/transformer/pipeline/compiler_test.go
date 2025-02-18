package pipeline

import (
	"io/ioutil"
	"testing"

	"github.com/gojek/merlin/pkg/transformer/types/series"
	"github.com/gojek/merlin/pkg/transformer/types/table"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"sigs.k8s.io/yaml"

	"github.com/gojek/merlin/pkg/transformer/feast"
	"github.com/gojek/merlin/pkg/transformer/spec"
	"github.com/gojek/merlin/pkg/transformer/symbol"
)

func TestCompiler_Compile(t *testing.T) {
	type (
		fields struct {
			sr           symbol.Registry
			feastClients feast.Clients
			feastOptions *feast.Options
			logger       *zap.Logger
		}

		want struct {
			expressions     []string
			jsonPaths       []string
			preloadedTables map[string]table.Table
			preprocessOps   []Op
			postprocessOps  []Op
		}
	)

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	tests := []struct {
		name             string
		fields           fields
		specYamlFilePath string
		want             want
		wantErr          bool
		expError         error
	}{
		{
			name: "preprocess input only",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_preprocess_input_only.yaml",
			want: want{
				expressions: []string{
					"Now()",
					"variable1",
				},
				jsonPaths: []string{
					"$.entity_1[*].id",
					"$.entity_2.id",
					"$.entity_3",
					"$.entity_2",
				},
				preloadedTables: map[string]table.Table{
					"preload_1": *table.New(
						series.New([]string{"Apple", "Banana", "Zara", "Sandra", "Merlion"}, series.String, "First Name"),
						series.New([]string{"Cider", "Man", "Vuitton", "Zawaska", "Krabby"}, series.String, "Last Name"),
						series.New([]int{25, 18, 35, 32, 23}, series.Int, "Age"),
						series.New([]float64{48.8, 68, 75, 55, 57.22}, series.Float, "Weight"),
						series.New([]bool{true, false, true, false, false}, series.Bool, "Is VIP"),
					),
				},
				preprocessOps: []Op{
					&FeastOp{},
					&CreateTableOp{},
					&VariableDeclarationOp{},
				},
				postprocessOps: []Op{},
			},
			wantErr: false,
		},
		{
			name: "no pipeline",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_no_pipeline.yaml",
			want: want{
				expressions:     []string{},
				jsonPaths:       []string{},
				preloadedTables: map[string]table.Table{},
				preprocessOps:   []Op{},
				postprocessOps:  []Op{},
			},
			wantErr: false,
		},
		{
			name: "preprocess postprocess input only",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_input_only.yaml",
			want: want{
				expressions: []string{
					"Now()",
					"variable1",
				},
				jsonPaths: []string{
					"$.entity_1[*].id",
					"$.entity_2.id",
					"$.entity_3",
					"$.entity_2",
				},
				preloadedTables: map[string]table.Table{
					"preload_1": *table.New(
						series.New([]string{"Apple", "Banana", "Zara", "Sandra", "Merlion"}, series.String, "First Name"),
						series.New([]string{"Cider", "Man", "Vuitton", "Zawaska", "Krabby"}, series.String, "Last Name"),
						series.New([]int{25, 18, 35, 32, 23}, series.Int, "Age"),
						series.New([]float64{48.8, 68, 75, 55, 57.22}, series.Float, "Weight"),
						series.New([]bool{true, false, true, false, false}, series.Bool, "Is VIP"),
					),
					"preload_2": *table.New(
						series.New([]string{"Apple", "Banana", "Zara", "Sandra", "Merlion"}, series.String, "First Name"),
						series.New([]string{"Cider", "Man", "Vuitton", "Zawaska", "Krabby"}, series.String, "Last Name"),
						series.New([]int{25, 18, 35, 32, 23}, series.Int, "Age"),
						series.New([]float64{48.8, 68, 75, 55, 57.22}, series.Float, "Weight"),
						series.New([]string{"TRUE", "FALSE", "TRUE", "FALSE", "FALSE"}, series.String, "Is VIP"),
					),
				},
				preprocessOps: []Op{
					&FeastOp{},
					&CreateTableOp{},
					&VariableDeclarationOp{},
				},
				postprocessOps: []Op{
					&FeastOp{},
					&CreateTableOp{},
					&VariableDeclarationOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "preprocess with input and transformation",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_preprocess_input_and_transformation.yaml",
			want: want{
				expressions: []string{
					"Now()",
					"variable1",
					"entity_2_table.Col('col1')",
					"Now().Hour()",
					"transformed_entity_3_table.Col('col5').Mean()",
					"transformed_entity_3_table.Col('col5').Max()",
				},
				jsonPaths: []string{
					"$.entity_1[*].id",
					"$.entity_2.id",
					"$.entity_3",
					"$.entity_2",
				},
				preprocessOps: []Op{
					&FeastOp{},
					&CreateTableOp{},
					&VariableDeclarationOp{},
					&EncoderOp{},
					&TableTransformOp{},
					&TableJoinOp{},
					&VariableDeclarationOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "sequential table dependency",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_sequential_table_dependency.yaml",
			want: want{
				expressions: []string{
					"Now()",
					"variable1",
					"entity_2_table.Col('col1')",
					"Now().Hour()",
				},
				jsonPaths: []string{
					"$.entity_1[*].id",
					"$.entity_2.id",
					"$.entity_3",
					"$.entity_2",
				},
				preprocessOps: []Op{
					&FeastOp{},
					&CreateTableOp{},
					&VariableDeclarationOp{},
					&TableTransformOp{},
					&TableTransformOp{},
					&TableJoinOp{},
					&TableJoinOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "preprocess - conditional, filter and slice row - valid",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
			},
			specYamlFilePath: "./testdata/valid_table_transform_conditional_filtering.yaml",
			want: want{
				expressions: []string{
					`zero`,
					`driver_table.Col("rating") * 2 <= 7`,
					`driver_table.Col("rating") * 1`,
					`driver_table.Col("rating") * 2 >= 8`,
					`driver_table.Col("rating") * 1.5`,
					`transformed_driver_table.Col('driver_performa').Max()`,
				},
				jsonPaths: []string{
					"$.customer.id",
					"$.drivers[*]",
				},
				preprocessOps: []Op{
					&VariableDeclarationOp{},
					&CreateTableOp{},
					&TableTransformOp{},
					&VariableDeclarationOp{},
					&JsonOutputOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "preprocess - postprocess input and output - valid",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/input_output.yaml",
			want: want{
				expressions: []string{
					"Now()",
					"variable1",
				},
				jsonPaths: []string{
					"$.entity_2.id",
					"$.entity_3",
					"$.entity_2",
					"$.path_1",
					"$.path_2",
				},
				preprocessOps: []Op{
					&CreateTableOp{},
					&VariableDeclarationOp{},
					&EncoderOp{},
					&JsonOutputOp{},
				},
				postprocessOps: []Op{
					&CreateTableOp{},
					&VariableDeclarationOp{},
					&EncoderOp{},
					&JsonOutputOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "feast preprocess",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_feast_preprocess.yaml",
			want: want{
				expressions: []string{
					"customer_id",
				},
				jsonPaths: []string{
					"$.customer.id",
					"$.drivers[*]",
					"$.drivers[*].id",
				},
				preprocessOps: []Op{
					&VariableDeclarationOp{},
					&CreateTableOp{},
					&FeastOp{},
					&TableTransformOp{},
					&TableJoinOp{},
					&TableTransformOp{},
					&JsonOutputOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "feast with expression",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/valid_feast_expression.yaml",
			want: want{
				expressions: []string{
					"customer_level",
					"customer_id",
					"Now().Hour()",
				},
				jsonPaths: []string{
					"$.customer.id",
					"$.drivers[*]",
					"$.drivers[*].id",
				},
				preprocessOps: []Op{
					&VariableDeclarationOp{},
					&CreateTableOp{},
					&FeastOp{},
					&TableTransformOp{},
					&TableJoinOp{},
					&TableTransformOp{},
					&JsonOutputOp{},
				},
			},
			wantErr: false,
		},
		{
			name: "preprocess - postprocess input and output - invalid",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/invalid_output.yaml",
			wantErr:          true,
			expError:         errors.New("unable to compile preprocessing pipeline: variable entity_5_table is not registered"),
		},
		{
			name: "invalid scale column - standard scale has zero standard deviation",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/invalid_standard_scale_column.yaml",
			wantErr:          true,
			expError:         errors.New("unable to compile preprocessing pipeline: standard scaler require non zero standard deviation"),
		},
		{
			name: "invalid scale column - min max scale has same min and max",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/invalid_min_max_scale_column.yaml",
			wantErr:          true,
			expError:         errors.New("unable to compile preprocessing pipeline: minmax scaler require different value between min and max"),
		},
		{
			name: "invalid encode column - encoder specified is not exist",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/invalid_encode_column.yaml",
			wantErr:          true,
			expError:         errors.New("unable to compile preprocessing pipeline: unknown name notFoundEncoder (1:1)\n | notFoundEncoder\n | ^"),
		},
		{
			name: "invalid - variables that want to store is not exist",
			fields: fields{
				sr:           symbol.NewRegistry(),
				feastClients: feast.Clients{},
				feastOptions: &feast.Options{
					CacheEnabled:  true,
					CacheSizeInMB: 100,
				},
				logger: logger,
			},
			specYamlFilePath: "./testdata/invalid_variables_in_transformations.yaml",
			wantErr:          true,
			expError:         errors.New("unable to compile preprocessing pipeline: unknown name transformed_entity_100_table (1:1)\n | transformed_entity_100_table.Col('col5').Mean()\n | ^"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Compiler{
				sr:           tt.fields.sr,
				feastClients: tt.fields.feastClients,
				feastOptions: tt.fields.feastOptions,
				logger:       tt.fields.logger,
			}

			yamlBytes, err := ioutil.ReadFile(tt.specYamlFilePath)
			assert.NoError(t, err)

			jsonBytes, err := yaml.YAMLToJSON(yamlBytes)
			assert.NoError(t, err)

			var stdSpec spec.StandardTransformerConfig
			err = protojson.Unmarshal(jsonBytes, &stdSpec)
			assert.NoError(t, err)

			got, err := c.Compile(&stdSpec)
			if tt.wantErr {
				assert.EqualError(t, err, tt.expError.Error())
				return
			}
			assert.NoError(t, err)

			for _, jsonPath := range tt.want.jsonPaths {
				assert.NotNil(t, got.compiledJsonpath.Get(jsonPath), "json path not compiled", jsonPath)
			}

			for _, expression := range tt.want.expressions {
				assert.NotNil(t, got.compiledExpression.Get(expression), "expression not compiled", expression)
			}

			assert.Equal(t, len(tt.want.preloadedTables), len(got.preloadedTables))
			for i, tb := range tt.want.preloadedTables {
				assert.Equal(t, tb, got.preloadedTables[i])
			}

			assert.Equal(t, len(tt.want.preprocessOps), len(got.preprocessOps))
			for i, op := range tt.want.preprocessOps {
				assert.IsType(t, op, got.preprocessOps[i], "different type")
			}

			assert.Equal(t, len(tt.want.postprocessOps), len(got.postprocessOps))
			for i, op := range tt.want.postprocessOps {
				assert.IsType(t, op, got.postprocessOps[i], "different type")
			}
		})
	}
}
