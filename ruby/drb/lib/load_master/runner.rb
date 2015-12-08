require "yaml"
require "load_master/slaves"
require "load_master/test"

module LoadMaster
  class Runner
    attr_reader :test, :slaves_uris

    def initialize(url, requests, concurrency, config_file)
      @test = Test.new(url, requests, concurrency)
      @slaves_uris = YAML.load_file(config_file)
    end

    def run
      p slaves.load(test)
    end

    private

    def slaves
      @slaves ||= Slaves.new(slaves_uris)
    end
  end
end
