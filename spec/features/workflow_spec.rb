%w( capybara capybara/rspec capybara-webkit ).each do |lib|
  require lib
end

include Capybara::DSL

describe 'visiting the homepage' do
  before do
    Capybara.configure do |config|
      config.run_server = false
      config.current_driver = :webkit
      config.app_host = 'http://localhost:8080'
      config.javascript_driver = :webkit
    end
  end

  it 'allows the user to visit the homepage' do
    visit '/'

    expect(page).to have_content 'Hello world'
  end

  it 'allows the user to see a list of todo items' do
    visit '/'

    expect(page).to have_css 'div#todo_items'
    expect(page).to have_content 'Cat'
    expect(page).to have_content 'Dog'
    expect(page).to have_content 'Giraffe'
    expect(page).to have_content 'Zebra'
    expect(page).to have_content 'Elephant'
  end
end