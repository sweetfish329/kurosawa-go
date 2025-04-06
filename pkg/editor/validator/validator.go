package validator

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

// カスタムエラー定義
var (
	ErrInvalidPath      = errors.New("invalid file path")
	ErrInvalidTimeRange = errors.New("invalid time range")
	ErrEmptyPath        = errors.New("empty path")
	ErrNotRegularFile   = errors.New("not a regular file")
)

type Validator interface {
	Validate() error
}

// ValidateFilePath はファイルパスの妥当性を検証します
func ValidateFilePath(path string) error {
	if path == "" {
		return ErrEmptyPath
	}

	// パスの正規化と絶対パス変換
	cleanPath := filepath.Clean(path)
	absPath, err := filepath.Abs(cleanPath)
	if err != nil {
		return err
	}

	// パスの存在確認
	fileInfo, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrInvalidPath
		}
		return err
	}

	// 通常ファイルであることを確認
	if !fileInfo.Mode().IsRegular() {
		return ErrNotRegularFile
	}

	return nil
}

// ValidateTimeRange は時間範囲の妥当性を検証します
func ValidateTimeRange(start, end time.Duration) error {
	if start < 0 || end < 0 {
		return ErrInvalidTimeRange
	}

	if end < start {
		return ErrInvalidTimeRange
	}

	return nil
}

// NewFileValidator ファイル関連の検証を行うための便利な構造体を返します
type FileValidator struct {
	Path string
}

func NewFileValidator(path string) *FileValidator {
	return &FileValidator{Path: path}
}

func (v *FileValidator) Validate() error {
	return ValidateFilePath(v.Path)
}

// NewTimeRangeValidator 時間範囲の検証を行うための便利な構造体を返します
type TimeRangeValidator struct {
	Start time.Duration
	End   time.Duration
}

func NewTimeRangeValidator(start, end time.Duration) *TimeRangeValidator {
	return &TimeRangeValidator{
		Start: start,
		End:   end,
	}
}

func (v *TimeRangeValidator) Validate() error {
	return ValidateTimeRange(v.Start, v.End)
}
