import static Services.getPolicies
import static Services.waitForDeployment
import static Services.waitForViolation
import spock.lang.Unroll
import objects.Deployment
import java.util.stream.Collectors

class SystemPoliciesTest extends  BaseSpecification {
    @Unroll
    def "Verify policy #policyname is triggered" (String policyname, Deployment deployment,
                                                  String testId, String deploymentName) {
        when:
        "Create a Deployment"
        orchestrator.createDeployment(deployment)

        and:
        "Validate if Violation and Deployment is present"
        assert getPolicies().stream()
                .filter { f -> f.getName() == policyname }
                .collect(Collectors.toList()).size() == 1

        assert waitForDeployment(deployment.getName())

        then:
        "Verify Violation #policyname is triggered"
        assert waitForViolation(deploymentName,  policyname, 10)

         cleanup:
        "Remove Deployment #deploymentName"
        orchestrator.deleteDeployment(deploymentName)

        where:
        "Data inputs are :"

        policyname | deployment | testId | deploymentName

        "Container Port 22" | new Deployment()
                .setName ("qaport22")
                .setImage ("nginx")
                .addPort (22)
                .addLabel ( "app", "test" ) | "C311" | "qaport22"

        "Apache Struts: CVE-2017-5638" | new Deployment()
                .setName ( "qacve" )
                .setImage ( "apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C938" | "qacve"

        "Heartbleed: CVE-2014-0160" | new Deployment()
                .setName ("qaheartbleed")
                .setImage ("apollo-dtr.rox.systems/legacy-apps/ssl-terminator")
                .addLabel ( "app", "test" ) | "C947" | "qaheartbleed"

        "Wget in Image" | new Deployment()
                .setName ("qawget")
                .setImage ("apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C939" | "qawget"

        "90-Day Image Age" | new Deployment()
                .setName ("qa90days" )
                .setImage ("apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ("app", "test" ) | "C810" | "qa90days"

        "Aptitude Package Manager (apt) in Image" | new Deployment()
                .setName ("qaapt" )
                .setImage ("apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C931" | "qaapt"

        "30-Day Scan Age" | new Deployment()
                .setName ( "qa30days" )
                .setImage ( "apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C941" | "qa30days"

        "Maximum CVSS >= 7" | new Deployment()
                .setName ( "qacvss" )
                .setImage ( "apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C933" | "qacvss"

        "Shellshock: CVE-2014-6271" | new Deployment()
                .setName ("qashellshock" )
                .setImage ("apollo-dtr.rox.systems/legacy-apps/ssl-terminator")
                .addLabel ( "app", "test" ) | "C948" | "qashellshock"

        "Curl in Image" | new Deployment()
                .setName ("qacurl")
                .setImage ("apollo-dtr.rox.systems/legacy-apps/struts-app:latest")
                .addLabel ( "app", "test" ) | "C948" | "qacurl"

        "DockerHub NGINX 1.10" | new Deployment()
                .setName ("qanginx")
                .setImage ("nginx:1.10")
                .addLabel ( "app", "test" ) | "C823" | "qanginx"

        "Any Vulnerabilities" | new Deployment()
                .setName ("qaanyvul" )
                .setImage ("apollo-dtr.rox.systems/legacy-apps/struts-app")
                .addLabel ( "app", "test" ) | "C814" | "qaanyvul"
    }
}
