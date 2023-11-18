/*
 * Swagger Petstore - OpenAPI 3.0
 *
 * This is a example for performance study based on the OpenAPI 3.0 specification.
 *
 * API version: 1.0.11
 * Contact: skalicky.martin@iotdomu.cz
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package database

import "github.com/uptrace/bun"

type RelItemCart struct {
	bun.BaseModel `bun:"table:rel_item_cart"`

	Id int64 `bun:"id,pk,autoincrement"`

	ItemSeqId int64

	CartSeqId int64
}
