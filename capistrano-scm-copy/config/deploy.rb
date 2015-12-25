# config valid only for current version of Capistrano
lock '3.4.0'

set :application, "example"
set :deploy_to, "~/example"
set :keep_releases, 5
set :log_level, :info

set :bundle_without, %i(development test)

# capistrano-scm-copyの設定
set :scm, :copy
