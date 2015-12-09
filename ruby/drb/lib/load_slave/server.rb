require "load_slave/loader"

module LoadSlave
  class Server
    def generate_loader(test)
      Loader.new(test)
    end
  end
end
