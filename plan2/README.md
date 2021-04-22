# plan2

## Dependency

- Go 1.16
- gqlgen
- chi
- sqlboiler

## Usage

```bash
# ローカルで起動
make run

# dockerで起動
docker-compose up -d
```

```graphql
query findUsers($first: Int, $after: String, $last: Int, $before: String, $orderBy: [UserOrder!]) {
    users(first: $first, after: $after, before: $before, last: $last, orderBy: $orderBy) {
      pageInfo {
        endCursor
        hasNextPage
        hasPreviousPage
        startCursor
      }
      edges{
        cursor
        node {
          ... on User {
            id
            name
          }
        }
      }
    	totalCount
    }
}

{
  "first": 5,
  "after": null,
	"before": null,
	"last": null,
  "orderBy": [{"field": "CREATED_AT", "direction": "DESC"}]
}
```

TODO: バッケージ整理

## 参考

https://github.com/sky0621/study-graphql/blob/v0.10.0/try01/src/backend/server.go
https://github.com/ctrlxc/graphql-example/blob/main/src/repository/repository.go
https://github.com/shufo/go-graphql-boilerplate/blob/master/resolver/user_resolver.go
https://github.com/graphql/graphiql/blob/39f5e45a9d2bb469260b62ff8057cc988d853470/packages/graphql-language-service-parser/benchmark/fixtures/github.graphql
https://github.com/jjjjackson/gqlgen-example/blob/master/presentation/schema/user.graphqls
https://github.com/kshamiev/sungora_history/blob/9434595e79630110f0c20dbdf73fd66a2f72bbad/pkg/app/database.go
