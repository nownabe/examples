require "drb"
require "parallel"

module LoadMaster
  class Slaves
    def initialize(drb_uris)
      @slaves = drb_uris.map { |uri| DRbObject.new_with_uri(uri) }
    end

    def load(test)
      Parallel.map(@slaves) do |slave|
        slave.generate_loader(test).execute
      end
    end
  end
end
