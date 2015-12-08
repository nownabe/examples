require "drb"
require "load_slave/server"

module LoadSlave
  class Runner
    attr_reader :host, :port

    def initialize(host: "localhost", port: 55331)
      @host = host
      @port = port
    end

    def run
      puts "Starting LoadSlave server"
      DRb.start_service(uri, Server.new)
      puts "Waiting requests"
      DRb.thread.join
    end

    def uri
      @uri ||= "druby://#{host}:#{port}"
    end
  end
end
