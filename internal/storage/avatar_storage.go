package storage

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	apperror "github.com/wasilisk/doit-api/internal/app_error"
)

type AvatarStorage struct {
	baseDir string
	baseURL string
}

func NewAvatarStorage(baseDir string, baseURL string) *AvatarStorage {
	return &AvatarStorage{baseDir: baseDir, baseURL: baseURL}
}

var allowedExtensions = map[string]bool{
	".jpg": true, ".jpeg": true, ".png": true, ".webp": true,
}

func (s *AvatarStorage) Validate(filename string) error {
	if !allowedExtensions[filepath.Ext(filename)] {
		return apperror.New(apperror.CodeFileTypeNotAllowed)
	}
	return nil
}

func (s *AvatarStorage) Save(file multipart.File, ext string) (string, error) {
	if err := os.MkdirAll(s.baseDir, os.ModePerm); err != nil {
		return "", apperror.New(apperror.CodeAvatarUploadFailed)
	}

	filename := uuid.New().String() + ext
	destPath := filepath.Join(s.baseDir, filename)

	dest, err := os.Create(destPath)
	if err != nil {
		return "", apperror.New(apperror.CodeAvatarUploadFailed)
	}
	defer dest.Close()

	if _, err := io.Copy(dest, file); err != nil {
		os.Remove(destPath)
		return "", apperror.New(apperror.CodeAvatarUploadFailed)
	}

	return filename, nil
}

func (s *AvatarStorage) Delete(avatarURL string) error {
	if avatarURL == "" {
		return nil
	}
	path := filepath.Join(s.baseDir, filepath.Base(avatarURL))
	if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	return nil
}

func (s *AvatarStorage) PublicURL(filename string) string {
	return s.baseURL + "/static/avatars/" + filename
}
