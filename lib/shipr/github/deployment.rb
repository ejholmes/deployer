module Shipr
  module GitHub
    class Deployment
      attr_reader :repo
      attr_reader :params

      def initialize(repo, params)
        @repo   = repo
        @params = params
      end

      def self.create(*args)
        new(*args).create
      end

      def create
        install_webhooks
        create_deployment
      end

    private

      def install_webhooks
        DeploymentHook.install(repo)
        DeploymentStatusHook.install(repo)
      end

      def create_deployment
        client.create_deployment repo, params
      end

      def client
        Shipr.github
      end
    end
  end
end
