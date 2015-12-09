require "drb"
require "load_master/test"

module LoadSlave
  class Loader
    include DRb::DRbUndumped

    AB_BINARY = "/usr/sbin/ab"

    attr_reader :test

    def initialize(test)
      @test = test
    end

    def execute
      puts "Execute `#{command}`"

      start_time = Time.now
      `#{command}`
      time = Time.now - start_time

      puts "Finished ab (#{time}s)"
      time
    end

    private

    def command
      "#{AB_BINARY} -n #{test.requests} -c #{test.concurrency} #{test.url}"
    end
  end
end
