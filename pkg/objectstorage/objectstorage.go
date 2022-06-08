/*
 *     Copyright 2022 The Dragonfly Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate mockgen -package mocks -source objectstorage.go -destination ./mocks/objectstorage_mock.go

package objectstorage

import (
	"fmt"
	"io"
	"time"
)

type ObjectMetadata struct {
	// Key is object key
	Key string

	// ContentDisposition is Content-Disposition header
	ContentDisposition string

	// ContentEncoding is Content-Encoding header
	ContentEncoding string

	// ContentLanguage is Content-Language header
	ContentLanguage string

	// ContentLanguage is Content-Length header
	ContentLength int64

	// ContentType is Content-Type header
	ContentType string

	// Etag is Etag header
	Etag string
}

type BucketMetadata struct {
	// Name is bucket name
	Name string

	// CreateAt is bucket create time.
	CreateAt time.Time
}

type ObjectStorage interface {
	// GetBucketMetadata returns metadata of bucket.
	GetBucketMetadata(bucketName string) (*BucketMetadata, error)

	// CreateBucket creates bucket of object storage.
	CreateBucket(bucketName string) error

	// DeleteBucket deletes bucket of object storage.
	DeleteBucket(bucketName string) error

	// ListBucketMetadatas returns metadata of buckets.
	ListBucketMetadatas() ([]*BucketMetadata, error)

	// GetObjectMetadata returns metadata of object.
	GetObjectMetadata(bucketName, objectKey string) (*ObjectMetadata, error)

	// GetOject returns data of object.
	GetOject(bucketName, objectKey string) (io.ReadCloser, error)

	// CreateObject creates data of object.
	CreateObject(bucketName, objectKey string, reader io.Reader) error

	// DeleteObject deletes data of object.
	DeleteObject(bucketName, objectKey string) error

	// ListObjectMetadatas returns metadata of objects.
	ListObjectMetadatas(bucketName, prefix, marker string, limit int64) ([]*ObjectMetadata, error)
}

// New object storage interface.
func New(name, region, endpoint, accessKey, secretKey string) (ObjectStorage, error) {
	switch name {
	case ServiceNameS3:
		return newS3(region, endpoint, accessKey, secretKey)
	case ServiceNameOSS:
		return newOSS(region, endpoint, accessKey, secretKey)
	}

	return nil, fmt.Errorf("unknow service name %s", name)
}