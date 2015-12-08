require "drb"

module LoadSlave
  class Loader
    include DRb::DRbUndumped

    AB_BINARY = "/usr/sbin/ab"

    attr_reader :url, :requests, :concurrency

    def initialize(url, requests, concurrency)
      @url         = url
      @requests    = requests
      @concurrency = concurrency
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
      "#{AB_BINARY} -n #{requests} -c #{concurrency} #{url}"
    end
  end
end
