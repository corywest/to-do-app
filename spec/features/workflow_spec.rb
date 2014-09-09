%w( capybara capybara/rspec capybara-webkit ).each do |lib|
  require lib
end

include Capybara::DSL

Capybara.configure do |config|
  config.run_server = false
  config.current_driver = :webkit
  config.app_host = 'http://localhost:8080'
  config.javascript_driver = :webkit
end

describe 'visiting the homepage' do
  it 'allows the user to visit the homepage' do
    visit '/'

    expect(page).to have_content "Hello world"
  end
end