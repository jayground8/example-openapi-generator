# Example Code

openapi-generator-cli를 통해서 Client 코드를 생성하는 예제

- Naver Cloud CDN API를 OpenAPI 3.0 Specification 맞춰서 작성하여 Client 코드를 생성한다.
- Credentials 정보와 함께 정상 작동 하는 것을 확인한다.


Docker로 CLI 실행

```bash
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i /local/openapi.yml \
    -g go \
    --additional-properties=withGoMod=false \
    -o /local/client
```
생성된 코드 중 `client.go` 파일에 Naver Cloud 인증을 위해서 아래의 코드를 추가한다.

```go
import (
	"crypto"
	...생략

	"github.com/jayground8/example-openapi-generator/ncloud/credentials"
	"github.com/jayground8/example-openapi-generator/hmac"
)
```

```go
queryString := ""
if len(url.RawQuery) > 0 {
  queryString = "?" + url.RawQuery
}

if auth := credentials.LoadCredentials(credentials.DefaultCredentialsChain()); auth != nil {
  timestamp := strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
  signer := hmac.NewSigner(auth.SecretKey(), crypto.SHA256)
  signature, _ := signer.Sign(method, path+queryString, auth.AccessKey(), timestamp)

  localVarRequest.Header.Add("x-ncp-apigw-timestamp", timestamp)
  localVarRequest.Header.Add("x-ncp-iam-access-key", auth.AccessKey())
  localVarRequest.Header.Add("x-ncp-apigw-signature-v1", signature)
}
```