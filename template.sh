export project_name=vblog
export package_name=github.com/vcycyv
export entity_capital=Blog
export entity=blog
export entity_plural=blogs
export entity_capital_plural=Blogs

mv "models/entity.go" "models/${entity}.go"
mv "routers/entity.go" "routers/${entity}.go"

grep -rl '{project_name}' * | xargs -i@ sed -i "s/{project_name}/$project_name/g" @
grep -rl '{package_name}' * | xargs -i@ sed -i "s~{package_name}~$package_name~g" @
grep -rl '{entity_capital}' * | xargs -i@ sed -i "s/{entity_capital}/$entity_capital/g" @
grep -rl '{entity}' * | xargs -i@ sed -i "s/{entity}/$entity/g" @
grep -rl '{entity_plural}' * | xargs -i@ sed -i "s/{entity_plural}/$entity_plural/g" @
grep -rl '{entity_capital_plural}' * | xargs -i@ sed -i "s/{entity_capital_plural}/$entity_capital_plural/g" @
