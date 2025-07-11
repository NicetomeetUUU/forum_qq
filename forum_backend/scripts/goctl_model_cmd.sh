goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/user \
    -c \
    -t user \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/admin \
    -c \
    -t admin \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/category \
    -c \
    -t category \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/post \
    -c \
    -t post \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/comment \
    -c \
    -t comment \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/user_like \
    -c \
    -t user_like \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/user_follow \
    -c \
    -t user_follow \
    --style goZero

goctl model mysql datasource \
    --url 'root:root542@tcp(127.0.0.1:3307)/qq_forum' \
    -d ../app/forum/model/post_category \
    -c \
    -t post_category \
    --style goZero



