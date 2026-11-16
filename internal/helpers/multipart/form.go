package multipart

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/textproto"
	"testing"

	"github.com/stretchr/testify/require"
)

type Form struct {
	buffer *bytes.Buffer
	writer *multipart.Writer
}

// NewForm создаёт пустую форму multipart/form-data.
func NewForm() *Form {
	buffer := &bytes.Buffer{}
	return &Form{
		buffer: buffer,
		writer: multipart.NewWriter(buffer),
	}
}

// AddField добавляет в форму текстовое поле.
func (f *Form) AddField(t *testing.T, name, value string) {
	t.Helper()
	require.NoErrorf(t, f.writer.WriteField(name, value), "Не удалось добавить поле %q в multipart-форму", name)
}

// AddFile добавляет в форму файл с указанным типом содержимого.
func (f *Form) AddFile(t *testing.T, fieldName, fileName, contentType string, content []byte) {
	t.Helper()

	header := make(textproto.MIMEHeader)
	header.Set("Content-Disposition", fmt.Sprintf(`form-data; name=%q; filename=%q`, fieldName, fileName))
	header.Set("Content-Type", contentType)

	part, err := f.writer.CreatePart(header)
	require.NoErrorf(t, err, "Не удалось создать часть для файла %q", fileName)

	_, err = part.Write(content)
	require.NoErrorf(t, err, "Не удалось записать файл %q в multipart-форму", fileName)
}

// Build закрывает форму и возвращает тело запроса и его Content-Type.
func (f *Form) Build(t *testing.T) (body string, contentType string) {
	t.Helper()
	require.NoError(t, f.writer.Close(), "Не удалось завершить multipart-форму")

	return f.buffer.String(), f.writer.FormDataContentType()
}
