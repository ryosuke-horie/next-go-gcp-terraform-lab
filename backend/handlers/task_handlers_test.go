package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/models"
	"github.com/ryosuke-horie/next-go-gcp-terraform-lab/repository"
)

// TestCreateTaskHandler は CreateTaskHandler のテスト
func TestCreateTaskHandler(t *testing.T) {
	// モックリポジトリの作成
	mockRepo := &repository.MockTaskRepository{
		CreateTaskFunc: func(ctx context.Context, task *models.Task) error {
			// テスト用にIDを設定
			task.ID = 1
			// 作成日時を設定（任意）
			task.CreatedAt = time.Now()
			return nil
		},
	}

	// ハンドラの初期化
	taskHandler := NewTaskHandler(mockRepo)

	// テストケースの定義
	tests := []struct {
		name           string
		input          map[string]string
		setupMock      func()
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "正常系",
			input: map[string]string{
				"title":  "正常系のタスク",
				"detail": "詳細タスク",
			},
			setupMock: func() {
				mockRepo.CreateTaskFunc = func(ctx context.Context, task *models.Task) error {
					task.ID = 1
					return nil
				}
			},
			expectedStatus: http.StatusCreated,
			expectedBody: map[string]interface{}{
				"id":           float64(1),
				"title":        "正常系のタスク",
				"detail":       "詳細タスク",
				"is_completed": false,
			},
		},
		{
			name:           "Invalid Input",
			input:          nil, // 空のリクエストボディ
			setupMock:      func() {},
			expectedStatus: http.StatusBadRequest,
			expectedBody:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// モックの設定
			tt.setupMock()

			// リクエストボディの準備
			var body bytes.Buffer
			if tt.input != nil {
				if err := json.NewEncoder(&body).Encode(tt.input); err != nil {
					t.Fatalf("Failed to encode input: %v", err)
				}
			}

			// /taskにPostリクエストを送る
			req := httptest.NewRequest(http.MethodPost, "/task", &body)
			req.Header.Set("Content-Type", "application/json")

			// レスポンス記録用のRecorderを作成
			rr := httptest.NewRecorder()

			// ハンドラの呼び出し
			taskHandler.CreateTaskHandler(rr, req)

			// ステータスコードの検証
			if rr.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, rr.Code)
			}

			// レスポンスボディの検証（期待される場合）
			if tt.expectedBody != nil {
				var respBody map[string]interface{}
				if err := json.NewDecoder(rr.Body).Decode(&respBody); err != nil {
					t.Fatalf("Failed to decode response body: %v", err)
				}

				for key, expectedValue := range tt.expectedBody {
					if respBody[key] != expectedValue {
						t.Errorf("Expected %s to be %v, got %v", key, expectedValue, respBody[key])
					}
				}

				// created_at の存在確認
				if _, ok := respBody["created_at"]; !ok {
					t.Errorf("Expected created_at field")
				}
			}
		})
	}
}
