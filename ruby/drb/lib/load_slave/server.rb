require "load_slave/loader"

module LoadSlave
  class Server
    def generate_loader(url, requests, concurrency)
      Loader.new(url, requests, concurrency)
    end
  end
end
